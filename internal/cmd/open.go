package cmd

import (
	"bufio"

	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"development-environment-cli/internal/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(openCmd)
}

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "进入当前目录的开发模式",
	Run: func(cmd *cobra.Command, args []string) {
		handleOpenCommand()
	},
}

func handleOpenCommand() {

	utils.EnsureMacOS()

	utils.EnsureCommand("code")
	utils.EnsureCommand("git")
	utils.EnsureCommand("gh")
	utils.EnsureCommand("docker")
	utils.EnsureCommand("mkdir")

	openCmd := flag.NewFlagSet("open", flag.ExitOnError)
	openCmd.Parse(os.Args[2:])

	if openCmd.NArg() != 1 {
		fmt.Println("open 命令有且只能有一个目录参数")
	}

	dir := openCmd.Arg(0)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("[ERROR] 目录 %s 不存在，是否创建？(Y/n): ", dir)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" || strings.ToLower(input) == "yes" || strings.ToLower(input) == "y" {
			fmt.Printf("[INFO] 正在创建目录 %s...\n", dir)
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				fmt.Printf("[ERROR] 创建目录 %s 失败: %v\n", dir, err)
				return
			}
			fmt.Printf("[INFO] 目录 %s 创建成功\n", dir)

			if err := os.Chdir(dir); err != nil {
				fmt.Printf("[ERROR] 切换到目录 %s 失败: %v\n", dir, err)
				return
			}

			cmd := exec.Command("git", "init")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Printf("[ERROR] 在目录 %s 初始化 git 仓库失败: %v\n", dir, err)
				return
			}
			fmt.Printf("[INFO] 在目录 %s 初始化 git 仓库成功\n", dir)
		} else {
			fmt.Println("[INFO] 操作已取消")
			return
		}
	}

	// 判断dir里是否包含.git目录
	if _, err := os.Stat(filepath.Join(dir, ".git")); os.IsNotExist(err) {
		fmt.Printf("[ERROR] 目录 %s 不是一个 Git 仓库\n", dir)
		return
	}

	// 显示当前目录
	if cwd, err := os.Getwd(); err == nil {
		fmt.Printf("当前目录: %s\n", cwd)
	} else {
		fmt.Printf("[ERROR] 获取当前目录失败: %v\n", err)
	}

	// 切换到指定目录
	if err := os.Chdir(dir); err != nil {
		fmt.Printf("[ERROR] 切换到目录 %s 失败: %v\n", dir, err)
		return
	}

	// 显示切换后的目录
	if cwd, err := os.Getwd(); err == nil {
		fmt.Printf("切换后的目录: %s\n", cwd)
	} else {
		fmt.Printf("[ERROR] 获取切换后的目录失败: %v\n", err)
	}

	// 执行 `code .` 命令
	cmd := exec.Command("code", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("[ERROR] 执行 `code .` 失败: %v\n", err)
		return
	}

}
