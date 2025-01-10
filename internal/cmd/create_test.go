package cmd

import (
	"testing"

	"development-environment-cli/internal/utils"

	"github.com/spf13/cobra"
)

func TestHandleCreateCommand(t *testing.T) {
	// 创建一个新的 root 命令
	rootCmd := &cobra.Command{Use: "root"}
	rootCmd.AddCommand(createCommand)

	// 测试用例
	// tests := []utils.CommandTestCase{
	// 	{[]string{"create", "myproject"}, true},
	// 	{[]string{"create"}, false},
	// }
	tests := []utils.CommandTestCase{
		{Args: []string{"create", "myproject"}, ExpectSuccess: true},
		{Args: []string{"create"}, ExpectSuccess: false},
		{Args: []string{"create", "arg1", "args2"}, ExpectSuccess: false},
		{Args: []string{"create", "arg1", "args2", "arg3"}, ExpectSuccess: false},
	}

	utils.RunCommandTests(t, rootCmd, tests)
}
