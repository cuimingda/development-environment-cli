package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
)

const version = "1.0.0"

func main() {

	debugOsArgs()
	fatalIfNotExistDir(".git")
	fatalIfNotExistDir(".vscode")

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "doctor":
			HandleDoctorCommand()
			return
		case "web":
			HandleWebCommand(os.Args[2:])
			return
		default:
			message := fmt.Sprintf("[ERROR] 未知命令 - %s", os.Args[1])
			colorizeMessage := colorize(message, "red")
			log.Fatalln(colorizeMessage)
			return
		}
	}

	handleDevCommand()
}

func debugOsArgs() {
	log.Printf("typeof of os.Args - %s\n", reflect.TypeOf(os.Args))

	log.Printf("len(os.Args) - %d", len(os.Args))
	for i, arg := range os.Args {
		log.Printf("os.Args[%d] - %s", i, arg)
	}
}

func fatalIfNotExistDir(dir string) {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		message := fmt.Sprintf("[ERROR] 当前目录没有 %s 目录，终止执行", dir)
		colorizeMessage := colorize(message, "red")
		log.Fatalln(colorizeMessage)
	}
}

func getWorkingDirOrFatal() string {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("获取当前工作目录失败: %v", err)
	}
	log.Printf("当前工作目录: %s", workingDir)

	return workingDir
}

func getWorkingDirName() string {
	workingDir := getWorkingDirOrFatal()

	// 获取当前工作目录的名称，类似于 shell 中的 basename
	dirName := filepath.Base(workingDir)
	log.Printf("当前目录名称: %s", dirName)

	return dirName
}

func handleDevCommand() {

	workingDir := getWorkingDirOrFatal()
	dirName := getWorkingDirName()

	cmd := exec.Command(
		"docker", "run",
		"-it",
		"--rm",
		"-w", "/app",
		"-v", workingDir+":/app",
		"--name", dirName+"-container",
		"cuimingda/development-environment",
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("运行 Alpine Linux 容器失败: %v", err)
	}
}

func HandleWebCommand(args []string) {

	webCmd := flag.NewFlagSet("web", flag.ExitOnError)
	port := webCmd.String("port", "5173", "端口号")

	webCmd.Parse(args)

	workingDir := getWorkingDirOrFatal()
	dirName := getWorkingDirName()

	log.Printf("端口号: %s", *port)
	cmd := exec.Command(
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

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("运行 Alpine Linux 容器失败: %v", err)
	}
}

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

func HandleVersionFlag() {
	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "显示版本信息")
	flag.Parse()

	if showVersion {
		fmt.Println("development-environment-cli 版本:", version)
		os.Exit(0)
	}
}
