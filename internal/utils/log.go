package utils

import (
	"fmt"
	"log"
	"os"
)

func printMessageLog(message string) {
	log.Println(message)
}

func printErrorLog(err error) {
	log.Printf("错误: %v", err)
}

func fatalError(err error, message string) {
	if err != nil {
		printMessageLog(message)
		printErrorLog(err)
		exitWithError()
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
