package cmd

import (
	"development-environment-cli/internal/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dev",
	Short: "a command line tool for developing in docker container",
}

func init() {

	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().BoolVarP(&utils.Verbose, "verbose", "v", false, "verbose output 1")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
