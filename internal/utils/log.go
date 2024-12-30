package utils

import (
	"fmt"
	"log"
	"os"
)

var Verbose bool

func PrintActionLog(message string) {
	if Verbose {
		log.Println(message)
	}
}

func PrintInfoLog(message string) {
	if Verbose {
		log.Println(message)
	}
}

func PrintSuccessLog(message string) {
	if Verbose {
		log.Println(message)
	}
}

func PrintErrorLog(message string) {
	if Verbose {
		log.Println(message)
	}
}

func PrintFormatLog(format string, args ...any) {
	if Verbose {
		message := fmt.Sprintf(format, args...)
		log.Println(message)
	}
}

func PrintInfoMessage(message string) {
	fmt.Println(message)
}

func PrintSuccessMessage(message string) {
	fmt.Println(message)
}

func PrintErrorMessage(message string) {
	fmt.Println(message)
}

func PrintFormatMessage(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	fmt.Println(message)
}

func FatalWithFormatMessage(format string, args ...any) {
	color := "\033[31m"
	reset := "\033[0m"
	message := fmt.Sprintf(format, args...)
	fmt.Printf("[FATAL] %s%s%s\n", color, message, reset)
	os.Exit(1)
}

func fatalError(err error, message string) {
	if err != nil {
		PrintFormatMessage("[FATAL] %s: %v", message, err)
		exitWithError()
	}
}
