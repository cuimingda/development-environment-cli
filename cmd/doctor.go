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
	checkCommand("go1", "version")
}

func checkCommand(name string, arg string) {
	cmd := exec.Command(name, arg)
	output, err := cmd.Output()
	if err != nil {
		printError("%s未安装或无法执行%s %s命令", name, name, arg)
	} else {
		printSuccess("%s版本: %s", name, output)
	}
}

func printError(format string, args ...any) {
	printColored("red", format, args...)
}

func printSuccess(format string, args ...any) {
	printColored("green", format, args...)
}

func printColored(color string, format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	colorizeMessage := colorize(message, color)
	fmt.Println(colorizeMessage)
}

func colorize(message string, color string) string {
	var colorCode string

	switch color {
	case "red":
		colorCode = "\033[31m"
	case "green":
		colorCode = "\033[32m"
	default:
		colorCode = ""
	}

	if colorCode == "" {
		return message
	}

	return fmt.Sprintf("%s%s\033[0m", colorCode, message)
}
