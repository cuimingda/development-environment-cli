// 执行外部命令
package utils

import (
	"os"
	"os/exec"
)

func ExecuteCommandWithOutput(name string, arg ...string) string {
	EnsureCommand(name)

	cmd := exec.Command(name, arg...)
	output, err := cmd.Output()
	FatalOnError(err, "executing command %s failed", name)

	return TrimmedStringFromBytes(output)
}

func ExecuteCommand(name string, args ...string) {
	EnsureCommand(name)

	cmd := exec.Command(name, args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	FatalOnError(err, "executing command %s failed", name)
}
