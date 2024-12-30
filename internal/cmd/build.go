package cmd

import (
	"development-environment-cli/internal/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var buildCommand = &cobra.Command{
	Use:   "build",
	Short: "build to ./bin/dev",
	Run: func(cmd *cobra.Command, args []string) {
		handleBuildCommand()
	},
}
var DisableLocalBinary bool
var DisableDockerImage bool

func init() {
	rootCmd.AddCommand(buildCommand)
	buildCommand.Flags().BoolVarP(&DisableLocalBinary, "disable-local-binary", "", false, "not to build to local bin")
	buildCommand.Flags().BoolVarP(&DisableDockerImage, "disable-docker-image", "", false, "not to build docker image")
}

func handleBuildCommand() {

	utils.PrintVerboseMessage("开始构建")

	if utils.Verbose {
		fmt.Printf("DisableLocalBinary: %v\n", DisableLocalBinary)
		fmt.Printf("DisableDockerImage: %v\n", DisableDockerImage)
	}

	utils.EnsurePath(".git")
	utils.EnsurePath("Dockerfile")
	utils.EnsureGitRemoteOrigin("git@github.com:cuimingda/development-environment-cli.git")

	if !DisableLocalBinary {
		utils.EnsurePath("main.go")
		utils.ExecuteCommand("go", "build", "-o", "bin/dev", ".")
	} else {
		fmt.Println("跳过构建本地二进制文件")
	}

	if !DisableDockerImage {
		utils.ExecuteCommand("docker", "build", "--tag", "cuimingda/development-environment:latest", ".")
	} else {
		fmt.Println("跳过构建 Docker 镜像")
	}

	utils.PrintVerboseMessage("构建完成")
}
