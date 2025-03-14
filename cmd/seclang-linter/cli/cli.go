package cli

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(runCmd)
}

const cliDescription = `
An extensible linter built to find sytnax errors in 
Coraza's specific implementation of SecLang.
`

var rootCmd = &cobra.Command{
	Use:   "seclang-linter [command]",
	Short: "seclang-linter finds errors in Coraza's SecLang syntax",
	Long:  cliDescription,
}

func Execute() error {
	return rootCmd.Execute()
}
