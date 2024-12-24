package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func HandleWebCommand(args []string) {

	webCmd := flag.NewFlagSet("web", flag.ExitOnError)
	port := webCmd.String("port", "", "端口号")

	webCmd.Parse(args)

	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("获取当前工作目录失败: %v", err)
	}
	log.Printf("当前工作目录: %s", workingDir)

	// 获取当前工作目录的名称，类似于 shell 中的 basename
	dirName := filepath.Base(workingDir)
	log.Printf("当前目录名称: %s", dirName)

	var cmd *exec.Cmd

	if *port == "" {
		log.Println("端口号: 未设置")
		cmd = exec.Command("docker", "run", "-it", "--rm", "-w", "/app", "-v", workingDir+":/app", "--name", dirName+"-container", "cuimingda/development-environment")
	} else {
		log.Printf("端口号: %s", *port)
		cmd = exec.Command("docker", "run", "-it", "--rm", "-w", "/app", "-v", workingDir+":/app", "--publish", *port+":"+*port, "--network", "local-network", "--name", dirName+"-container", "cuimingda/development-environment")
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatalf("运行 Alpine Linux 容器失败: %v", err)
	}
}
