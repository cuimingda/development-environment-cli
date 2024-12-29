package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {

	// if isDebug() {
	// 	debugOsArgs()
	// }

	var rootCmd = &cobra.Command{
		Use:   "dev",
		Short: "这是一个示例应用程序",
		Long:  `这是一个使用 Cobra 库创建的示例应用程序。`,
	}

	var versionsCmd = &cobra.Command{
		Use:   "versions",
		Short: "显示版本信息",
		Run:   HandleVersionsCommand,
	}

	var openCmd = &cobra.Command{
		Use:   "open",
		Short: "打开项目",
		Run: func(cmd *cobra.Command, args []string) {
			handleOpenCommand()
		},
	}

	var webCmd = &cobra.Command{
		Use:   "web",
		Short: "启动Web服务",
		Run: func(cmd *cobra.Command, args []string) {
			HandleWebCommand(os.Args[2:])
		},
	}

	var buildSelfCmd = &cobra.Command{
		Use:   "build-self",
		Short: "构建自身",
		Run: func(cmd *cobra.Command, args []string) {
			BuildSelfCommand()
		},
	}

	var devCmd = &cobra.Command{
		Use:   ".",
		Short: "启动开发环境",
		Run: func(cmd *cobra.Command, args []string) {
			handleDevCommand()
		},
	}

	rootCmd.AddCommand(
		devCmd,
		versionsCmd,
		openCmd,
		webCmd,
		buildSelfCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func isVerbose() bool {
	for _, arg := range os.Args {
		if arg == "-V" {
			return true
		}
		lowerArg := strings.ToLower(arg)
		if lowerArg == "--verbose" || lowerArg == "-verbose" {
			return true
		}
	}
	return false
}

func isDebug() bool {
	for _, arg := range os.Args {
		if arg == "-D" {
			return true
		}
		lowerArg := strings.ToLower(arg)
		if lowerArg == "--debug" || lowerArg == "-debug" {
			return true
		}
	}
	return false
}
