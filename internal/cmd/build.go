package cmd

import (
	"development-environment-cli/internal/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(buildCommand)
}

var buildCommand = &cobra.Command{
	Use:   "build",
	Short: "build to ./bin/dev",
	Run: func(cmd *cobra.Command, args []string) {
		handleBuildCommand()
	},
}

func handleBuildCommand() {

	printVerboseMessage("开始构建")

	utils.EnsureMacOS()
	utils.EnsurePath(".git")
	utils.EnsureCommand("go")

	utils.ExecuteCommand("go", "build", "-o", "bin/dev", ".")

	printVerboseMessage("构建完成")
}
