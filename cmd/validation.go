package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ensure(condition bool, format string, args ...any) {
	if !condition {
		fmt.Printf(format, args...)
		exitWithError()
	}
}

func isCommandAvailable(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

func ensureCommand(command string) {
	ensure(
		isCommandAvailable(command),
		"[ERROR] %s 未安装",
		command,
	)
}

func isPathAvailable(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func ensurePath(path string) {

	ensure(
		isPathAvailable(path),
		"[ERROR] 文件或目录 %s 不存在",
		path,
	)
}

func isAlpine() bool {
	return isPathAvailable("/etc/alpine-release")
}

func ensureAlpine() {
	ensure(
		isAlpine(),
		"[ERROR] 当前系统不是 %s",
		"Alpine Linux",
	)
}

func isMacOS() bool {
	return runtime.GOOS == "darwin"
}

func ensureMacOS() {
	ensure(
		isMacOS(),
		"[ERROR] 当前系统不是 %s",
		"MacOS",
	)
}

func isGitRemoteOrigin(expectedURL string) bool {
	remoteURL := getCommandOutputOrFatal("git", "config", "--get", "remote.origin.url")
	return remoteURL == expectedURL
}

func ensureGitRemoteOrigin(url string) {
	ensure(
		isGitRemoteOrigin(url),
		"[ERROR] 当前git项目的remote origin不是 %s",
		url,
	)
}

func fatalWithoutCondition(condition bool, message string) {
	if !condition {
		printMessageLog(message)
		exitWithError()
	}
}

func printFormatLog(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	printMessageLog(message)
}
