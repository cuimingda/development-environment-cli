package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
)

const version = "1.0.0"

func main() {

	if isDebug() {
		debugOsArgs()
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "doctor":
			HandleDoctorCommand()
			return
		case "open":
			handleOpenCommand()
			return
		case "web":
			HandleWebCommand(os.Args[2:])
			return
		default:
			fatalWithFormatMessage("[ERROR] 未知命令 - %s", os.Args[1])
			return
		}
	}

	handleDevCommand()
}

func isVerbose() bool {
	for _, arg := range os.Args {
		if arg == "-V" {
			return true
		}
		lowerArg := strings.ToLower(arg)
		if lowerArg == "--verbose" || lowerArg == "-verbose" {
			return true
		}
	}
	return false
}

func isDebug() bool {
	for _, arg := range os.Args {
		if arg == "-D" {
			return true
		}
		lowerArg := strings.ToLower(arg)
		if lowerArg == "--debug" || lowerArg == "-debug" {
			return true
		}
	}
	return false
}

func handleOpenCommand() {

	fatalIfNotExistCommand("code")
	fatalIfNotExistCommand("git")
	fatalIfNotExistCommand("gh")
	fatalIfNotExistCommand("docker")
	fatalIfNotExistCommand("mkdir")

	openCmd := flag.NewFlagSet("open", flag.ExitOnError)
	openCmd.Parse(os.Args[2:])

	if openCmd.NArg() != 1 {
		fatalWithMessage("open 命令有且只能有一个目录参数")
	}

	dir := openCmd.Arg(0)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("[ERROR] 目录 %s 不存在，是否创建？(Y/n): ", dir)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" || strings.ToLower(input) == "yes" || strings.ToLower(input) == "y" {
			fmt.Printf("[INFO] 正在创建目录 %s...\n", dir)
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				fmt.Printf("[ERROR] 创建目录 %s 失败: %v\n", dir, err)
				return
			}
			fmt.Printf("[INFO] 目录 %s 创建成功\n", dir)

			if err := os.Chdir(dir); err != nil {
				fmt.Printf("[ERROR] 切换到目录 %s 失败: %v\n", dir, err)
				return
			}

			cmd := exec.Command("git", "init")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Printf("[ERROR] 在目录 %s 初始化 git 仓库失败: %v\n", dir, err)
				return
			}
			fmt.Printf("[INFO] 在目录 %s 初始化 git 仓库成功\n", dir)
		} else {
			fmt.Println("[INFO] 操作已取消")
			return
		}
	}

	// 判断dir里是否包含.git目录
	if _, err := os.Stat(filepath.Join(dir, ".git")); os.IsNotExist(err) {
		fmt.Printf("[ERROR] 目录 %s 不是一个 Git 仓库\n", dir)
		return
	}

	// 显示当前目录
	if cwd, err := os.Getwd(); err == nil {
		fmt.Printf("当前目录: %s\n", cwd)
	} else {
		fmt.Printf("[ERROR] 获取当前目录失败: %v\n", err)
	}

	// 切换到指定目录
	if err := os.Chdir(dir); err != nil {
		fmt.Printf("[ERROR] 切换到目录 %s 失败: %v\n", dir, err)
		return
	}

	// 显示切换后的目录
	if cwd, err := os.Getwd(); err == nil {
		fmt.Printf("切换后的目录: %s\n", cwd)
	} else {
		fmt.Printf("[ERROR] 获取切换后的目录失败: %v\n", err)
	}

	// 执行 `code .` 命令
	cmd := exec.Command("code", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("[ERROR] 执行 `code .` 失败: %v\n", err)
		return
	}

}

// CheckCommandInstalled 检查命令是否已安装
func CheckCommandInstalled(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

func fatalIfNotExistCommand(command string) {
	if !CheckCommandInstalled(command) {
		fatalWithFormatMessage("[ERROR] %s 未安装或无法执行 %s 命令", command, command)
	}
}

func fatalWithFormatMessage(format string, args ...interface{}) {
	color := "\033[31m"
	reset := "\033[0m"
	message := fmt.Sprintf(format, args...)
	fmt.Printf("[FATAL] %s%s%s\n", color, message, reset)
	os.Exit(1)
}

func fatalWithMessage(message string) {
	color := "\033[31m"
	reset := "\033[0m"
	fmt.Printf("[FATAL] %s%s%s\n", color, message, reset)
	os.Exit(1)
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
		fatalWithFormatMessage("[ERROR] 当前目录没有 %s 目录，终止执行", dir)
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

	fatalIfNotExistDir(".git")
	fatalIfNotExistDir(".vscode")

	fatalIfNotExistCommand("code")
	fatalIfNotExistCommand("git")
	fatalIfNotExistCommand("gh")
	fatalIfNotExistCommand("docker")
	fatalIfNotExistCommand("mkdir")

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

func HandleWebCommand(args []string) {

	fatalIfNotExistCommand("code")
	fatalIfNotExistCommand("git")
	fatalIfNotExistCommand("gh")
	fatalIfNotExistCommand("docker")
	fatalIfNotExistCommand("mkdir")

	fatalIfNotExistDir(".git")
	fatalIfNotExistDir(".vscode")

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
		exitSucceed()
	}
}
