package main

func BuildSelfCommand() {

	ensureMacOS()
	ensurePath(".git")
	ensurePath("Dockerfile")
	ensureGitRemoteOrigin("git@github.com:cuimingda/development-environment-cli.git")

	executeCommand("docker", "build", "--tag", "cuimingda/development-environment:latest", ".")
}
