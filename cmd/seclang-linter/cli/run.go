package cli

import (
	"fmt"

	"github.com/bak-minsu/seclang-linter/internal/parse"
	"github.com/spf13/cobra"
)

const runDescription = `
Run linters on files within given paths.
Given paths can be any of the following:
- Path to a file
- Path to a directory containing only seclang files
- Glob path, ex. "./some/path/*"
- Triple dot path, ex. "./..."
`

var runCmd = &cobra.Command{
	Use:   "run [OPTIONS] <path to seclang file> <additional paths to seclang files>...",
	Short: "Runs linter on given paths",
	Long:  runDescription,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := parse.ParseFile(args[0]); err != nil {
			fmt.Println(err)
		}
	},
}
