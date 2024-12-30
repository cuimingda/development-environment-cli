package utils

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
)

func GetWorkingDirOrFatal() string {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("获取当前工作目录失败: %v", err)
	}
	log.Printf("当前工作目录: %s", workingDir)

	return workingDir
}

func GetWorkingDirName() string {
	workingDir := GetWorkingDirOrFatal()

	// 获取当前工作目录的名称，类似于 shell 中的 basename
	dirName := filepath.Base(workingDir)
	log.Printf("当前目录名称: %s", dirName)

	return dirName
}

func GetCommandOutputOrFatal(name string, arg ...string) string {

	EnsureCommand(name)

	cmd := exec.Command(name, arg...)
	output, err := cmd.Output()
	if err != nil {
		fatalWithFormatMessage("Error getting %s output: %v", name, err)
	}

	return strings.TrimSpace(string(output))
}

func ExecuteCommand(name string, args ...string) {
	EnsureCommand(name)

	cmd := exec.Command(name, args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	fatalError(err, "执行命令失败")
}

// 终止命令
func exitWithError() {
	os.Exit(1)
}

func exitSucceed() {
	os.Exit(0)
}

func debugOsArgs() {
	log.Printf("typeof of os.Args - %s\n", reflect.TypeOf(os.Args))

	log.Printf("len(os.Args) - %d", len(os.Args))
	for i, arg := range os.Args {
		log.Printf("os.Args[%d] - %s", i, arg)
	}
}
