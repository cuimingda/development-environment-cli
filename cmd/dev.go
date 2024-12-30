package cmd

import (
	"development-environment-cli/internal/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(devCmd)
}

var devCmd = &cobra.Command{
	Use:   ".",
	Short: "进入当前目录的开发模式",
	Run: func(cmd *cobra.Command, args []string) {
		handleDevCommand()
	},
}

func handleDevCommand() {

	utils.EnsureMacOS()

	utils.EnsurePath(".git")
	utils.EnsurePath(".vscode")

	utils.EnsureCommand("code")
	utils.EnsureCommand("git")
	utils.EnsureCommand("gh")
	utils.EnsureCommand("docker")
	utils.EnsureCommand("mkdir")

	workingDir := utils.GetWorkingDirOrFatal()
	dirName := utils.GetWorkingDirName()

	utils.ExecuteCommand(
		"docker", "run",
		"-it",
		"--rm",
		"-w", "/app",
		"-v", workingDir+":/app",
		"--name", dirName+"-container",
		"cuimingda/development-environment",
	)
}
