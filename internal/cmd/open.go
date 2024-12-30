package cmd

import (
	"bufio"

	"fmt"
	"os"
	"strings"

	"development-environment-cli/internal/utils"

	"github.com/spf13/cobra"
)

var openCommand = &cobra.Command{
	Use:   "open [dir]",
	Short: "用vscode打开指定目录",
	Args:  cobra.ExactArgs(1),
	Run:   handleOpenCommand,
}

func init() {
	rootCommand.AddCommand(openCommand)
}

func confirm(format string, args ...any) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(format+"(Y/n)", args...)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	yes := input == "" || strings.ToLower(input) == "yes" || strings.ToLower(input) == "y"
	return yes
}

func handleOpenCommand(cmd *cobra.Command, args []string) {

	utils.EnsureMacOS()

	utils.EnsureCommand("git")
	utils.EnsureCommand("gh")
	utils.EnsureCommand("docker")

	dir := args[0]
	utils.PrintInfoLog("dir: %s", dir)
	utils.FatalIf(dir == "", "目录参数不能为空")

	utils.PrintInfoLog("IsPathAvailable: %v", utils.IsPathAvailable(dir))
	if !utils.IsPathAvailable(dir) {
		yes := confirm("目录 %s 不存在，是否创建？ ", dir)
		utils.FatalIfNot(yes, "操作已取消")
		utils.PrintInfoLog("yes: %v", yes)
		utils.CreateDir(dir)
	}

	utils.EnterDir(dir)
	utils.ExecuteCommand("code", ".")
}
