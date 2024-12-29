package main

import (
	"os/exec"
	"strings"
)

func getCommandOutputOrFatal(name string, arg ...string) string {

	fatalIfNotExistCommand(name)

	cmd := exec.Command(name, arg...)
	output, err := cmd.Output()
	if err != nil {
		fatalWithFormatMessage("Error getting %s output: %v", name, err)
	}

	return strings.TrimSpace(string(output))
}

func fatalIfNotGitRemote(url string) {
	remoteURL := getCommandOutputOrFatal("git", "config", "--get", "remote.origin.url")
	fatalWithoutCondition(
		remoteURL == url,
		"当前不是"+url,
	)
}

func BuildSelfCommand() {

	fatalIfNotMacOS()
	fatalIfNotExistPath(".git")
	fatalIfNotExistPath("Dockerfile")
	fatalIfNotGitRemote("git@github.com:cuimingda/development-environment-cli.git")

	executeCommand("docker", "build", "--tag", "cuimingda/development-environment:latest", ".")
}
