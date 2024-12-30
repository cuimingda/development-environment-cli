package cmd

import (
	"development-environment-cli/internal/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(buildSelfCommand)
}

var buildSelfCommand = &cobra.Command{
	Use:   "build-self",
	Short: "编译docker",
	Run: func(cmd *cobra.Command, args []string) {
		handleBuildSelfCommand()
	},
}

func handleBuildSelfCommand() {

	utils.EnsureMacOS()
	utils.EnsurePath(".git")
	utils.EnsurePath("Dockerfile")
	utils.EnsureGitRemoteOrigin("git@github.com:cuimingda/development-environment-cli.git")

	utils.ExecuteCommand("docker", "build", "--tag", "cuimingda/development-environment:latest", ".")
}
