package cmd

import (
	"development-environment-cli/internal/utils"

	"github.com/spf13/cobra"
)

var buildCommand = &cobra.Command{
	Use:   "build",
	Short: "build to ./bin/dev",
	Run:   handleBuildCommand,
}
var DisableLocalBinary bool
var DisableDockerImage bool

func init() {
	rootCommand.AddCommand(buildCommand)
	buildCommand.Flags().BoolVar(&DisableLocalBinary, "disable-local-binary", false, "not to build to local bin")
	buildCommand.Flags().BoolVar(&DisableDockerImage, "disable-docker-image", false, "not to build docker image")
}

func handleBuildCommand(cmd *cobra.Command, args []string) {

	utils.PrintActionLog("开始构建")

	utils.PrintInfoLog("DisableLocalBinary: %v\n", DisableLocalBinary)
	utils.PrintInfoLog("DisableDockerImage: %v\n", DisableDockerImage)

	utils.EnsurePath(".git")
	utils.EnsurePath("Dockerfile")
	utils.EnsureGitRemoteOrigin("git@github.com:cuimingda/development-environment-cli.git")

	if !DisableLocalBinary {
		utils.EnsurePath("main.go")
		utils.ExecuteCommand("go", "build", "-o", "bin/dev", ".")
	} else {
		utils.PrintActionLog("跳过构建本地二进制文件")
	}

	if !DisableDockerImage {
		utils.ExecuteCommand("docker", "build", "--tag", "cuimingda/development-environment:latest", ".")
	} else {
		utils.PrintActionLog("跳过构建 Docker 镜像")

	}

	utils.PrintActionLog("构建完成")
}
