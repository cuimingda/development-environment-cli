package cmd

import (
	"development-environment-cli/internal/utils"

	"github.com/spf13/cobra"
)

var dotCommand = &cobra.Command{
	Use:   ".",
	Short: "进入当前目录的开发模式",
	Run:   handleDotCommand,
}

func init() {
	rootCommand.AddCommand(dotCommand)
}

func handleDotCommand(cmd *cobra.Command, args []string) {
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
