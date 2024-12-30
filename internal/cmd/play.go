package cmd

import (
	"development-environment-cli/internal/utils"

	"github.com/spf13/cobra"
)

var playCommand = &cobra.Command{
	Use:   "play",
	Short: "创建一个临时apline容器",
	Run:   handlePlayCommand,
}

var Image string

func init() {
	rootCommand.AddCommand(playCommand)
	playCommand.Flags().StringVar(&Image, "image", "alpine:3.21", "docker image to run")
}

func handlePlayCommand(cmd *cobra.Command, args []string) {

	utils.EnsureMacOS()

	utils.FatalIf(Image == "", "image is required")

	utils.ExecuteCommand(
		"docker", "run",
		"-it",
		"--rm",
		"-w", "/app",
		Image,
	)
}
