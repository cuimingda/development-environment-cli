package cmd

import (
	"fmt"
	"os/exec"
	"regexp"

	"development-environment-cli/internal/utils"

	"github.com/spf13/cobra"
)

var versionsCommand = &cobra.Command{
	Use:   "versions",
	Short: "显示所有安装包的版本",
	Run:   HandleVersionsCommand,
}

func init() {
	rootCommand.AddCommand(versionsCommand)
}

func HandleVersionsCommand(cmd *cobra.Command, args []string) {
	utils.PrintActionLog("检查命令...")

	checkCommand("docker", "--version")
	checkCommand("git", "--version")
	checkCommand("node", "--version")
	checkCommand("npm", "--version")
	checkCommand("curl", "--version")
	checkCommand("gh", "--version")
	checkCommand("go", "version")
	checkCommand("go1", "version")

	utils.PrintActionLog("检查完成")
}

func checkCommand(name string, arg string) {
	// 检查命令是否存在
	if _, err := exec.LookPath(name); err != nil {
		fmt.Printf("%s: 未安装", name)

		return
	}

	// 执行带有参数的命令
	cmd := exec.Command(name, arg)

	output, err := cmd.CombinedOutput()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			fmt.Printf("%s 命令执行失败，退出状态码: %d，错误信息: %s", name, exitError.ExitCode(), string(output))
		} else {
			fmt.Printf("%s 命令执行失败: %s", name, err)
		}

		return
	}

	// 使用正则表达式提取版本号
	re := regexp.MustCompile(`\d+\.\d+\.\d+`)

	version := re.FindString(string(output))
	if version == "" {
		version = string(output)
	}

	fmt.Printf("%s: %s", name, version)
}
