package output

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Format is the global output format: json (default), pretty, table
var Format = "json"

// FormatResponse formats an API response JSON string according to the current Format.
func FormatResponse(jsonStr string) string {
	switch Format {
	case "pretty":
		return prettyFormat(jsonStr)
	case "table":
		return tableFormat(jsonStr)
	default:
		return jsonStr
	}
}

// PrintJSON prints data as formatted JSON to stdout.
func PrintJSON(v interface{}) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error formatting output: %v\n", err)
		return
	}
	fmt.Println(string(data))
}

// PrintSuccess prints a success message.
func PrintSuccess(msg string) {
	fmt.Printf("✅ %s\n", msg)
}

// PrintError prints an error message to stderr.
func PrintError(msg string) {
	fmt.Fprintf(os.Stderr, "❌ %s\n", msg)
}

// PrintInfo prints an info message.
func PrintInfo(msg string) {
	fmt.Printf("ℹ️  %s\n", msg)
}

// ─── pretty format ───────────────────────────────────────────────

func prettyFormat(jsonStr string) string {
	var raw map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &raw); err != nil {
		return jsonStr
	}

	var sb strings.Builder

	// Header: traceId
	if traceID, ok := raw["traceId"]; ok {
		sb.WriteString(fmt.Sprintf("TraceID:  %v\n", traceID))
	}

	// Status line
	code, _ := raw["code"]
	msg, _ := raw["msg"]
	codeNum, _ := code.(float64)
	if codeNum == 0 {
		sb.WriteString(fmt.Sprintf("Status:   ✅ OK (%s)\n", msg))
	} else {
		sb.WriteString(fmt.Sprintf("Status:   ❌ FAILED (code=%v, %s)\n", code, msg))
	}

	// Data section
	data, hasData := raw["data"]
	if !hasData || data == nil {
		return sb.String()
	}

	sb.WriteString("\n")

	dataMap, isMap := data.(map[string]interface{})
	if !isMap {
		prettyData, _ := json.MarshalIndent(data, "", "  ")
		sb.WriteString(string(prettyData))
		sb.WriteString("\n")
		return sb.String()
	}

	// Summary fields
	summaryKeys := []string{"total", "totalAmount", "successAmount", "failAmount",
		"page", "pageSize", "balance", "giftMoney", "availableTimeAddOn",
		"taskId", "url", "status", "output"}
	printed := map[string]bool{}
	for _, key := range summaryKeys {
		if v, ok := dataMap[key]; ok {
			sb.WriteString(fmt.Sprintf("  %-22s %v\n", key+":", v))
			printed[key] = true
		}
	}

	// List/array sections
	listKeys := []string{"items", "list", "details", "successDetails", "failDetails", "ids"}
	for _, key := range listKeys {
		arr, ok := dataMap[key].([]interface{})
		if !ok || len(arr) == 0 {
			continue
		}
		printed[key] = true
		sb.WriteString(fmt.Sprintf("\n  📋 %s (%d entries):\n", key, len(arr)))
		for i, item := range arr {
			m, ok := item.(map[string]interface{})
			if !ok {
				sb.WriteString(fmt.Sprintf("    [%d] %v\n", i+1, item))
				continue
			}
			sb.WriteString(fmt.Sprintf("    [%d]\n", i+1))
			for k, v := range m {
				switch val := v.(type) {
				case map[string]interface{}:
					inner, _ := json.MarshalIndent(val, "        ", "  ")
					sb.WriteString(fmt.Sprintf("      %-18s %s\n", k+":", string(inner)))
				case []interface{}:
					inner, _ := json.Marshal(val)
					sb.WriteString(fmt.Sprintf("      %-18s %s\n", k+":", string(inner)))
				default:
					sb.WriteString(fmt.Sprintf("      %-18s %v\n", k+":", v))
				}
			}
		}
	}

	// Remaining fields
	for k, v := range dataMap {
		if printed[k] {
			continue
		}
		switch val := v.(type) {
		case []interface{}:
			if len(val) > 0 {
				sb.WriteString(fmt.Sprintf("\n  📋 %s (%d entries):\n", k, len(val)))
				for i, item := range val {
					sb.WriteString(fmt.Sprintf("    [%d] %v\n", i+1, item))
				}
			}
		case map[string]interface{}:
			inner, _ := json.MarshalIndent(val, "  ", "  ")
			sb.WriteString(fmt.Sprintf("\n  %s:\n  %s\n", k, string(inner)))
		default:
			sb.WriteString(fmt.Sprintf("  %-22s %v\n", k+":", v))
		}
	}

	return sb.String()
}

// ─── table format ────────────────────────────────────────────────

func tableFormat(jsonStr string) string {
	var raw map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &raw); err != nil {
		return jsonStr
	}

	var sb strings.Builder

	// Status line
	code, _ := raw["code"]
	msg, _ := raw["msg"]
	codeNum, _ := code.(float64)
	if codeNum == 0 {
		sb.WriteString(fmt.Sprintf("[OK] %s\n", msg))
	} else {
		sb.WriteString(fmt.Sprintf("[FAIL] code=%v %s\n", code, msg))
	}

	data, hasData := raw["data"]
	if !hasData || data == nil {
		return sb.String()
	}

	dataMap, isMap := data.(map[string]interface{})
	if !isMap {
		prettyData, _ := json.MarshalIndent(data, "", "  ")
		sb.WriteString(string(prettyData))
		sb.WriteString("\n")
		return sb.String()
	}

	// Find the main array to display as table
	tableKeys := []string{"items", "list", "details", "successDetails", "failDetails"}
	var tableArr []interface{}
	var tableKey string
	for _, key := range tableKeys {
		arr, ok := dataMap[key].([]interface{})
		if ok && len(arr) > 0 {
			tableArr = arr
			tableKey = key
			break
		}
	}

	// Print summary fields inline
	summaryKeys := []string{"total", "totalAmount", "successAmount", "failAmount", "page", "pageSize",
		"balance", "giftMoney", "availableTimeAddOn", "taskId", "url", "status", "output"}
	summaryParts := []string{}
	for _, key := range summaryKeys {
		if v, ok := dataMap[key]; ok {
			summaryParts = append(summaryParts, fmt.Sprintf("%s=%v", key, v))
		}
	}
	if len(summaryParts) > 0 {
		sb.WriteString(strings.Join(summaryParts, "  "))
		sb.WriteString("\n")
	}

	if tableArr == nil || len(tableArr) == 0 {
		// No table data, show remaining as key-value
		for k, v := range dataMap {
			found := false
			for _, sk := range summaryKeys {
				if k == sk {
					found = true
					break
				}
			}
			if found {
				continue
			}
			switch val := v.(type) {
			case []interface{}:
				if len(val) > 0 {
					sb.WriteString(fmt.Sprintf("\n%s:\n", k))
					inner, _ := json.MarshalIndent(val, "", "  ")
					sb.WriteString(string(inner))
					sb.WriteString("\n")
				}
			case map[string]interface{}:
				inner, _ := json.MarshalIndent(val, "", "  ")
				sb.WriteString(fmt.Sprintf("\n%s:\n%s\n", k, string(inner)))
			default:
				sb.WriteString(fmt.Sprintf("%s: %v\n", k, v))
			}
		}
		return sb.String()
	}

	sb.WriteString(fmt.Sprintf("\n%s (%d rows):\n", tableKey, len(tableArr)))
	sb.WriteString(renderTable(tableArr))

	// Show other arrays (like failDetails if items was the main one)
	for _, key := range tableKeys {
		if key == tableKey {
			continue
		}
		arr, ok := dataMap[key].([]interface{})
		if ok && len(arr) > 0 {
			sb.WriteString(fmt.Sprintf("\n%s (%d rows):\n", key, len(arr)))
			sb.WriteString(renderTable(arr))
		}
	}

	return sb.String()
}

func renderTable(arr []interface{}) string {
	if len(arr) == 0 {
		return ""
	}

	// Collect all keys preserving order from first item
	allKeys := []string{}
	keySet := map[string]bool{}
	for _, item := range arr {
		m, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		for k := range m {
			if !keySet[k] {
				keySet[k] = true
				allKeys = append(allKeys, k)
			}
		}
	}

	if len(allKeys) == 0 {
		data, _ := json.MarshalIndent(arr, "", "  ")
		return string(data) + "\n"
	}

	// Build rows as string values
	widths := make(map[string]int)
	for _, k := range allKeys {
		widths[k] = len(k)
	}

	rows := []map[string]string{}
	for _, item := range arr {
		m, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		row := map[string]string{}
		for _, k := range allKeys {
			v := m[k]
			var s string
			switch val := v.(type) {
			case nil:
				s = "-"
			case string:
				s = val
			case map[string]interface{}, []interface{}:
				b, _ := json.Marshal(val)
				s = string(b)
				if len(s) > 40 {
					s = s[:37] + "..."
				}
			default:
				s = fmt.Sprintf("%v", val)
			}
			row[k] = s
			if len(s) > widths[k] {
				widths[k] = len(s)
			}
		}
		rows = append(rows, row)
	}

	// Cap column widths at 40
	for k, w := range widths {
		if w > 40 {
			widths[k] = 40
		}
	}

	var sb strings.Builder

	// Header
	for i, k := range allKeys {
		if i > 0 {
			sb.WriteString(" │ ")
		}
		sb.WriteString(fmt.Sprintf("%-*s", widths[k], strings.ToUpper(k)))
	}
	sb.WriteString("\n")

	// Separator
	for i, k := range allKeys {
		if i > 0 {
			sb.WriteString("─┼─")
		}
		sb.WriteString(strings.Repeat("─", widths[k]))
	}
	sb.WriteString("\n")

	// Data rows
	for _, row := range rows {
		for i, k := range allKeys {
			if i > 0 {
				sb.WriteString(" │ ")
			}
			s := row[k]
			if len(s) > widths[k] {
				s = s[:widths[k]-3] + "..."
			}
			sb.WriteString(fmt.Sprintf("%-*s", widths[k], s))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
