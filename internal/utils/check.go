// 所有返回bool的检查类函数
package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func isCommandAvailable(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

func isPathAvailable(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func isAlpine() bool {
	return isPathAvailable("/etc/alpine-release")
}

func isMacOS() bool {
	return runtime.GOOS == "darwin"
}

func isGitRemoteOrigin(expectedURL string) bool {
	remoteURL := ExecuteCommandWithOutput("git", "config", "--get", "remote.origin.url")
	return remoteURL == expectedURL
}
