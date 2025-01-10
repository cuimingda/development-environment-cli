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

	if err := utils.EnsureCommandArgsLength(args, 1); err != nil {
		return err
	}

	fmt.Println("create command called")
	return nil
}
