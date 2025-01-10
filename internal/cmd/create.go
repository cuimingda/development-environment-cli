package cmd

import (
	"development-environment-cli/internal/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var createCommand = &cobra.Command{
	Use:   "create",
	Short: "创建新项目",
	RunE:  handleCreateCommand,
}

func init() {
	rootCommand.AddCommand(createCommand)
}

func handleCreateCommand(_ *cobra.Command, args []string) error {

	err := utils.EnsureCommandArgsLength(args, 1)

	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	fmt.Println("create command called")
	return nil
}
