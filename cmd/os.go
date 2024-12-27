package main

import (
	"os"
	"os/exec"
)

func executeCommand(name string, args ...string) {
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
