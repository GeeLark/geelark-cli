package mcpdocs

import (
	"encoding/json"
	"io/fs"
	"regexp"
	"sort"
	"strings"
)

// enumPairRe matches "0=started" style enum pairs inside a description.
var enumPairRe = regexp.MustCompile(`(\d+)\s*=\s*([^,|=]+)`)

// EnumHints returns enum hint lines for keys present in the JSON output,
// based on the reference doc of the given CLI command. It lets any LLM
// interpret numeric enum fields (e.g. status: 2) without reading the docs.
// Returns "" when no hints apply.
func (d *Docs) EnumHints(command, output string) string {
	path, ok := d.Resolve(command)
	if !ok {
		return ""
	}
	hints := d.enumsForDoc(path)
	if len(hints) == 0 {
		return ""
	}
	var v any
	if !json.Valid([]byte(output)) || json.Unmarshal([]byte(output), &v) != nil {
		return ""
	}
	keys := map[string]bool{}
	collectKeys(v, keys)

	var lines []string
	seen := map[string]bool{}
	for seg, ls := range hints {
		if !keys[seg] {
			continue
		}
		for _, l := range ls {
			if !seen[l] {
				seen[l] = true
				lines = append(lines, l)
			}
		}
	}
	if len(lines) == 0 {
		return ""
	}
	sort.Strings(lines)
	return strings.Join(lines, "\n")
}

// enumsForDoc parses a reference doc and maps the last segment of each
// enum field path (e.g. "items[].status" -> "status") to its hint line.
// Only response-field tables (3 columns: Field | Type | Description) are
// considered; flag tables (2 columns) and request-param tables (4 columns)
// are skipped. Descriptions must contain at least two "N=value" pairs.
func (d *Docs) enumsForDoc(path string) map[string][]string {
	d.enumMu.Lock()
	defer d.enumMu.Unlock()
	if c, ok := d.enumCache[path]; ok {
		return c
	}
	m := map[string][]string{}
	b, err := fs.ReadFile(d.fsys, path)
	if err != nil {
		d.enumCache[path] = m
		return m
	}
	for _, line := range strings.Split(string(b), "\n") {
		cells := splitTableRow(line)
		if len(cells) != 3 {
			continue
		}
		field, desc := strings.TrimSpace(cells[0]), cells[2]
		if !strings.HasPrefix(field, "`") || !strings.HasSuffix(field, "`") {
			continue
		}
		field = strings.Trim(field, "`")
		if field == "" || strings.HasPrefix(field, "--") || strings.ContainsAny(field, " <") {
			continue
		}
		pairs := enumPairRe.FindAllStringSubmatch(desc, -1)
		if len(pairs) < 2 {
			continue
		}
		parts := make([]string, 0, len(pairs))
		for _, p := range pairs {
			parts = append(parts, strings.TrimSpace(p[1])+"="+strings.TrimSpace(p[2]))
		}
		seg := field
		if i := strings.LastIndex(seg, "."); i >= 0 {
			seg = seg[i+1:]
		}
		seg = strings.TrimSuffix(seg, "[]")
		m[seg] = append(m[seg], field+": "+strings.Join(parts, ", "))
	}
	d.enumCache[path] = m
	return m
}

// splitTableRow splits a markdown table row into trimmed cells.
// Returns nil if the line is not a table row.
func splitTableRow(line string) []string {
	line = strings.TrimSpace(line)
	if !strings.HasPrefix(line, "|") || !strings.HasSuffix(line, "|") {
		return nil
	}
	inner := strings.TrimSuffix(strings.TrimPrefix(line, "|"), "|")
	cells := strings.Split(inner, "|")
	for i := range cells {
		cells[i] = strings.TrimSpace(cells[i])
	}
	return cells
}

// collectKeys gathers all object keys from a decoded JSON value.
func collectKeys(v any, out map[string]bool) {
	switch t := v.(type) {
	case map[string]any:
		for k, val := range t {
			out[k] = true
			collectKeys(val, out)
		}
	case []any:
		for _, e := range t {
			collectKeys(e, out)
		}
	}
}
