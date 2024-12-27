package main

import (
	"log"
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
