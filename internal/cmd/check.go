package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"development-environment-cli/internal/utils"

	"github.com/spf13/cobra"
)

var checkStatusCommand = &cobra.Command{
	Use:   "check-status [dir]",
	Short: "检查指定目录下所有项目的git状态",
	Run: func(cmd *cobra.Command, args []string) {
		utils.FatalIf(len(args) < 1, "Usage: dev check-status <dir>")
		handleCheckStatusCommand(cmd, args)
	},
}

func init() {
	rootCommand.AddCommand(checkStatusCommand)
}

func getSubdirectories(dir string) ([]string, error) {
	var subdirs []string

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			subdirs = append(subdirs, filepath.Join(dir, file.Name()))
		}
	}

	return subdirs, nil
}

func checkGitStatus(dir string) (string, error) {
	cmd := exec.Command("git", "-C", dir, "status", "--porcelain")

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	if len(output) == 0 {
		return "Clean", nil
	}

	return "Uncommitted changes", nil
}

func checkGitPushStatus(dir string) (string, error) {
	cmd := exec.Command("git", "-C", dir, "rev-list", "@{u}..HEAD")

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	if len(output) == 0 {
		return "All changes pushed", nil
	}

	return "Unpushed changes", nil
}

func handleCheckStatusCommand(_ *cobra.Command, args []string) {
	utils.EnsureMacOS()

	dir := args[0]
	utils.PrintInfoLog("dir: %s", dir)
	utils.FatalIf(dir == "", "目录参数不能为空")

	utils.EnsurePath(dir)

	subdirs, err := getSubdirectories(dir)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	for _, subdir := range subdirs {
		fmt.Println("Directory:", subdir)

		status, err := checkGitStatus(subdir)
		if err != nil {
			fmt.Println("Error checking git status:", err)

			continue
		}

		fmt.Println("Git Status:", status)

		pushStatus, err := checkGitPushStatus(subdir)
		if err != nil {
			fmt.Println("Error checking git push status:", err)

			continue
		}

		fmt.Println("Git Push Status:", pushStatus)
		fmt.Println()
	}
}
