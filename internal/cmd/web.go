package cmd

import (
	"fmt"

	"development-environment-cli/internal/utils"

	"github.com/spf13/cobra"
)

var port int

var webCommand = &cobra.Command{
	Use:   "web",
	Short: "进入当前目录的web开发模式",
	Run:   handleWebCommand,
}

func init() {
	rootCommand.AddCommand(webCommand)
	webCommand.Flags().IntVarP(&port, "port", "p", 5173, "Web服务监听的端口号")
}

func handleWebCommand(cmd *cobra.Command, args []string) {

	utils.EnsureMacOS()

	utils.EnsureCommand("code")
	utils.EnsureCommand("git")
	utils.EnsureCommand("gh")
	utils.EnsureCommand("docker")
	utils.EnsureCommand("mkdir")

	utils.EnsurePath(".git")
	utils.EnsurePath(".vscode")

	workingDir := utils.GetWorkingDirOrFatal()
	dirName := utils.GetWorkingDirName()

	publishConfig := fmt.Sprintf("%d:%d", port, port)
	utils.ExecuteCommand(
		"docker", "run",
		"-it",
		"--rm",
		"-w", "/app",
		"-v", workingDir+":/app",
		"--publish", publishConfig,
		"--network", "local-network",
		"--name", dirName+"-container",
		"cuimingda/development-environment",
	)
}
