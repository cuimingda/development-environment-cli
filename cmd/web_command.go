package main

import (
	"flag"
	"log"
)

func HandleWebCommand(args []string) {

	ensureMacOS()

	ensureCommand("code")
	ensureCommand("git")
	ensureCommand("gh")
	ensureCommand("docker")
	ensureCommand("mkdir")

	ensurePath(".git")
	ensurePath(".vscode")

	webCmd := flag.NewFlagSet("web", flag.ExitOnError)
	port := webCmd.String("port", "5173", "端口号")

	webCmd.Parse(args)

	workingDir := getWorkingDirOrFatal()
	dirName := getWorkingDirName()

	log.Printf("端口号: %s", *port)
	executeCommand(
		"docker", "run",
		"-it",
		"--rm",
		"-w", "/app",
		"-v", workingDir+":/app",
		"--publish", *port+":"+*port,
		"--network", "local-network",
		"--name", dirName+"-container",
		"cuimingda/development-environment",
	)
}
