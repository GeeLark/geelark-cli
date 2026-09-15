// Package mcpdocs indexes the skills reference documentation and looks it up by
// CLI command path (e.g. "browser create" or "phone automation tiktok-login").
//
// Lookup keys are taken from each document's first heading (e.g.
// "# phone automation tiktok-login"), which is the canonical command path —
// no guessing from file names.
package mcpdocs

import (
	"fmt"
	"io/fs"
	"sort"
	"strings"
	"sync"
)

// Docs indexes one skills documentation tree.
type Docs struct {
	fsys fs.FS

	indexOnce sync.Once
	// lookup maps a CLI command path (e.g. "phone automation tiktok-login")
	// to its reference file. Keys come from the documents' first headings.
	lookup map[string]string

	enumMu sync.Mutex
	// enumCache maps a doc path to its last-key-segment -> hint lines index.
	enumCache map[string]map[string][]string
}

// New returns a Docs backed by fsys, which must contain the skills tree.
func New(fsys fs.FS) *Docs {
	return &Docs{
		fsys:      fsys,
		enumCache: map[string]map[string][]string{},
	}
}

func (d *Docs) index() map[string]string {
	d.indexOnce.Do(func() {
		lookup := make(map[string]string)
		_ = fs.WalkDir(d.fsys, ".", func(path string, entry fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if entry.IsDir() || !strings.HasSuffix(path, ".md") {
				return nil
			}
			b, err := fs.ReadFile(d.fsys, path)
			if err != nil {
				return err
			}
			if key := firstHeading(string(b)); key != "" {
				lookup[key] = path
			}
			return nil
		})
		d.lookup = lookup
	})
	return d.lookup
}

// firstHeading returns the text of the first "# " heading, trimmed of the
// leading "geelark-cli" if present. Returns "" if the document has none.
func firstHeading(doc string) string {
	for _, line := range strings.Split(doc, "\n") {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "# ") {
			continue
		}
		h := strings.TrimSpace(strings.TrimPrefix(line, "# "))
		return strings.TrimPrefix(h, "geelark-cli ")
	}
	return ""
}

// Get returns the reference document for a CLI command path
// (e.g. "browser create", "phone automation tiktok-login").
// The path may omit leading groups ("tiktok-login" resolves to
// "phone automation tiktok-login") — the lookup falls back to suffix matching.
// When nothing matches, the error lists the closest known command paths so the
// caller can retry with a full one.
func (d *Docs) Get(command string) (string, error) {
	if p, ok := d.Resolve(command); ok {
		b, err := fs.ReadFile(d.fsys, p)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	cmd := normalize(command)
	if cmd == "" {
		return "", fmt.Errorf("empty command")
	}
	if s := d.suggest(cmd); len(s) > 0 {
		return "", fmt.Errorf("no reference doc found for command %q; did you mean: %s", cmd, strings.Join(s, ", "))
	}
	return "", fmt.Errorf("no reference doc found for command %q", cmd)
}

// suggest returns up to five known command paths that contain every segment of
// cmd in order, e.g. "phone tiktok-login" suggests "phone automation
// tiktok-login". Requiring all segments keeps a command like "phone app list"
// from proposing every unrelated "... list" there is.
func (d *Docs) suggest(cmd string) []string {
	want := strings.Fields(cmd)
	if len(want) == 0 {
		return nil
	}

	var out []string
	for k := range d.index() {
		if containsInOrder(strings.Fields(k), want) {
			out = append(out, k)
		}
	}
	sort.Strings(out)
	if len(out) > 5 {
		out = out[:5]
	}
	return out
}

// containsInOrder reports whether want appears in segments as a subsequence.
func containsInOrder(segments, want []string) bool {
	i := 0
	for _, s := range segments {
		if i < len(want) && s == want[i] {
			i++
		}
	}
	return i == len(want)
}

// Resolve maps a CLI command path to its reference file path.
func (d *Docs) Resolve(command string) (string, bool) {
	lookup := d.index()

	cmd := normalize(command)
	if cmd == "" {
		return "", false
	}

	if p, ok := lookup[cmd]; ok {
		return p, true
	}

	// Suffix match: allow omitting leading groups, e.g. "tiktok-login" matches
	// "phone automation tiktok-login". A bare verb such as "list" matches many
	// keys; prefer the one closest to what was asked for (fewest segments) and
	// break ties alphabetically so the result is deterministic across runs.
	keys := make([]string, 0, len(lookup))
	for k := range lookup {
		if strings.HasSuffix(k, " "+cmd) {
			keys = append(keys, k)
		}
	}
	if len(keys) == 0 {
		return "", false
	}
	sort.Slice(keys, func(i, j int) bool {
		ni, nj := len(strings.Fields(keys[i])), len(strings.Fields(keys[j]))
		if ni != nj {
			return ni < nj
		}
		return keys[i] < keys[j]
	})
	return lookup[keys[0]], true
}

func normalize(command string) string {
	cmd := strings.TrimSpace(command)
	cmd = strings.TrimPrefix(cmd, "geelark-cli")
	return strings.TrimSpace(cmd)
}
