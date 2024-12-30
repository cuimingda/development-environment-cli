package cmd

import (
	"development-environment-cli/internal/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(buildDockerCommand)
}

var buildDockerCommand = &cobra.Command{
	Use:   "build-docker",
	Short: "编译docker",
	Run: func(cmd *cobra.Command, args []string) {
		handleBuildDockerCommand()
	},
}

func handleBuildDockerCommand() {

	utils.EnsureMacOS()
	utils.EnsurePath(".git")
	utils.EnsurePath("Dockerfile")
	utils.EnsureGitRemoteOrigin("git@github.com:cuimingda/development-environment-cli.git")

	utils.ExecuteCommand("docker", "build", "--tag", "cuimingda/development-environment:latest", ".")
	utils.ExecuteCommand("go", "build", "-o", "bin/dev", ".")
}
