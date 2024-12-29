package main

func handleDevCommand() {

	ensureMacOS()

	ensurePath(".git")
	ensurePath(".vscode")

	ensureCommand("code")
	ensureCommand("git")
	ensureCommand("gh")
	ensureCommand("docker")
	ensureCommand("mkdir")

	workingDir := getWorkingDirOrFatal()
	dirName := getWorkingDirName()

	executeCommand(
		"docker", "run",
		"-it",
		"--rm",
		"-w", "/app",
		"-v", workingDir+":/app",
		"--name", dirName+"-container",
		"cuimingda/development-environment",
	)
}
