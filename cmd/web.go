package cmd

import (
	"development-environment-cli/internal/utils"
	"flag"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(webCmd)
}

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "进入当前目录的web开发模式",
	Run: func(cmd *cobra.Command, args []string) {
		handleWebCommand(args)
	},
}

func handleWebCommand(args []string) {

	utils.EnsureMacOS()

	utils.EnsureCommand("code")
	utils.EnsureCommand("git")
	utils.EnsureCommand("gh")
	utils.EnsureCommand("docker")
	utils.EnsureCommand("mkdir")

	utils.EnsurePath(".git")
	utils.EnsurePath(".vscode")

	webCmd := flag.NewFlagSet("web", flag.ExitOnError)
	port := webCmd.String("port", "5173", "端口号")

	webCmd.Parse(args)

	workingDir := utils.GetWorkingDirOrFatal()
	dirName := utils.GetWorkingDirName()

	log.Printf("端口号: %s", *port)
	utils.ExecuteCommand(
		"docker", "run",
		"-it",
		"--rm",
		"-w", "/app",
		"-v", workingDir+":/app",
		"--publish", *port+":"+*port,
		"--network", "local-network",
		"--name", dirName+"-container",
		"cuimingda/development-environment",
	)
}
