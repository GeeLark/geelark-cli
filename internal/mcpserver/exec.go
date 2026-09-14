package mcpserver

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Outcome holds the result of one CLI invocation.
type Outcome struct {
	Stdout  string
	Stderr  string
	ExitErr error // non-nil when the process failed
}

// OK reports whether the command exited successfully.
func (o Outcome) OK() bool { return o.ExitErr == nil }

// FailureText returns the best error text: stderr first (cobra prints
// errors and usage there), then stdout, then the raw exit error.
func (o Outcome) FailureText() string {
	var b strings.Builder
	if o.Stderr != "" {
		b.WriteString(o.Stderr)
	}
	if o.Stdout != "" {
		if b.Len() > 0 {
			b.WriteString("\n")
		}
		b.WriteString(o.Stdout)
	}
	if b.Len() == 0 && o.ExitErr != nil {
		b.WriteString(o.ExitErr.Error())
	}
	return b.String()
}

// SelfBin returns the path of the running geelark-cli binary.
//
// Preference order:
//  1. GEELARK_CLI_PATH — tests / wrappers that replace argv[0]
//  2. os.Executable — absolute path; survives chdir unlike a relative argv[0]
//  3. os.Args[0] — last resort
func SelfBin() string {
	if p := os.Getenv("GEELARK_CLI_PATH"); p != "" {
		return p
	}
	if exe, err := os.Executable(); err == nil {
		if resolved, err := filepath.EvalSymlinks(exe); err == nil {
			return resolved
		}
		return exe
	}
	return os.Args[0]
}

// run executes: <self> <path...> <sub> <flags...> --format json
func (s *Server) run(ctx context.Context, path []string, sub string, params map[string]any) (Outcome, error) {
	flags, err := buildFlags(params)
	if err != nil {
		return Outcome{}, err
	}
	args := make([]string, 0, len(path)+1+len(flags)+2)
	args = append(args, path...)
	args = append(args, sub)
	args = append(args, flags...)
	args = append(args, "--format", "json")
	return execCLI(ctx, args), nil
}

// runHelp executes: <self> <path...> --help
func (s *Server) runHelp(ctx context.Context, path []string) Outcome {
	args := make([]string, 0, len(path)+1)
	args = append(args, path...)
	args = append(args, "--help")
	return execCLI(ctx, args)
}

// execCLI forks the same geelark-cli binary. Capturing the child process
// stdout/stderr keeps the MCP transport on the parent stdout untouched,
// so none of the CLI's fmt.Println sites need to change.
func execCLI(ctx context.Context, args []string) Outcome {
	cmd := exec.CommandContext(ctx, SelfBin(), args...)
	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return Outcome{Stdout: stdout.String(), Stderr: stderr.String(), ExitErr: err}
}

var flagNameRe = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9-]*$`)

// buildFlags converts a params map into pflag-style arguments.
// Keys become flag names (sorted for determinism); values are formatted
// as --key=value so string, bool, and slice flags all work.
func buildFlags(params map[string]any) ([]string, error) {
	if len(params) == 0 {
		return nil, nil
	}
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	out := make([]string, 0, len(params))
	for _, k := range keys {
		if !flagNameRe.MatchString(k) {
			return nil, fmt.Errorf("invalid flag name %q (must be letters, digits, dashes)", k)
		}
		s, err := formatValue(params[k])
		if err != nil {
			return nil, fmt.Errorf("invalid value for flag --%s: %w", k, err)
		}
		out = append(out, "--"+k+"="+s)
	}
	return out, nil
}

func formatValue(v any) (string, error) {
	switch t := v.(type) {
	case nil:
		return "", fmt.Errorf("null value")
	case string:
		return t, nil
	case bool:
		return strconv.FormatBool(t), nil
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64), nil
	case int:
		return strconv.Itoa(t), nil
	case int64:
		return strconv.FormatInt(t, 10), nil
	case []any:
		parts := make([]string, 0, len(t))
		for _, e := range t {
			s, err := formatValue(e)
			if err != nil {
				return "", err
			}
			parts = append(parts, s)
		}
		return strings.Join(parts, ","), nil
	default:
		return fmt.Sprintf("%v", t), nil
	}
}
