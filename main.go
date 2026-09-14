// geelark-cli — GeeLark CLI tool.
package main

import (
	"embed"
	"os"

	"github.com/geelark-tech/geelark-cli/cmd"
)

// skillsFS carries the reference documentation served by the mcp command.
// It is embedded here because go:embed cannot reach outside its own package
// directory, and skills/ is published from the repository root.
//
//go:embed all:skills
var skillsFS embed.FS

func main() {
	os.Exit(cmd.Execute(skillsFS))
}
