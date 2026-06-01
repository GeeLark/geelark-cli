// geelark-cli — GeeLark CLI tool.
package main

import (
	"os"

	"github.com/geelark-tech/geelark-cli/cmd"
)

func main() {
	os.Exit(cmd.Execute())
}
