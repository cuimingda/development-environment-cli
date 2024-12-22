package main

import (
	"fmt"
	"os/exec"
)

func HandleDoctorCommand() {

	// 检查命令
	checkCommand("docker", "--version")
	checkCommand("git", "--version")
	checkCommand("node", "--version")
	checkCommand("npm", "--version")
	checkCommand("curl", "--version")
	checkCommand("gh", "--version")
	checkCommand("go", "version")
}

func checkCommand(name string, arg string) {
	cmd := exec.Command(name, arg)
	output, err := cmd.Output()
	if err != nil {
		printColored(fmt.Sprintf("%s未安装或无法执行%s命令", name, name), "red")
	} else {
		printColored(fmt.Sprintf("%s版本: %s", name, output), "green")
	}
}

func printColored(message string, color string) {
	var colorCode string
	switch color {
	case "red":
		colorCode = "\033[31m"
	case "green":
		colorCode = "\033[32m"
	default:
		colorCode = "\033[0m"
	}
	fmt.Printf("%s%s\033[0m\n", colorCode, message)
}
