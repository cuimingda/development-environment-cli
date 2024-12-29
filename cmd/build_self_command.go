package main

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
