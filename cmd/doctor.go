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
		fmt.Printf("\033[31m%s未安装或无法执行%s命令\033[0m\n", name, name) // 红色
	} else {
		fmt.Printf("\033[32m%s版本: %s\033[0m\n", name, output) // 绿色
	}
}
