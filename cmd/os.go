package main

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func isAlpine() bool {
	if _, err := os.Stat("/etc/alpine-release"); err == nil {
		return true
	}
	return false
}

func isMacOS() bool {
	return runtime.GOOS == "darwin"
}

func fatalWithoutCondition(condition bool, message string) {
	if !condition {
		printMessageLog(message)
		exitWithError()
	}
}

func fatalIfNotAlpine() {
	fatalWithoutCondition(isAlpine(), "当前系统不是 Alpine Linux")
}

func fatalIfNotMacOS() {
	fatalWithoutCondition(isMacOS(), "当前系统不是 MacOS")
}

func getCommandOutputOrFatal(name string, arg ...string) string {

	fatalIfNotExistCommand(name)

	cmd := exec.Command(name, arg...)
	output, err := cmd.Output()
	if err != nil {
		fatalWithFormatMessage("Error getting %s output: %v", name, err)
	}

	return strings.TrimSpace(string(output))
}

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
