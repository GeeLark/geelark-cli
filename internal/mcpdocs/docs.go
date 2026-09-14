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
// The path may omit intermediate groups ("phone tiktok-login") — the lookup
// falls back to suffix matching.
func (d *Docs) Get(command string) (string, error) {
	if p, ok := d.Resolve(command); ok {
		b, err := fs.ReadFile(d.fsys, p)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	if normalize(command) == "" {
		return "", fmt.Errorf("empty command")
	}
	return "", fmt.Errorf("no reference doc found for command %q", normalize(command))
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

	// Suffix match: allow omitting intermediate groups, e.g. "phone tiktok-login"
	// matches "phone automation tiktok-login". Prefer the longest key.
	keys := make([]string, 0, len(lookup))
	for k := range lookup {
		if strings.HasSuffix(k, " "+cmd) {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys) // longest last
	if len(keys) > 0 {
		return lookup[keys[len(keys)-1]], true
	}
	return "", false
}

func normalize(command string) string {
	cmd := strings.TrimSpace(command)
	cmd = strings.TrimPrefix(cmd, "geelark-cli")
	return strings.TrimSpace(cmd)
}
