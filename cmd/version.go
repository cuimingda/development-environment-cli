package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "1.0.0"

func HandleVersionFlag() {
	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "显示版本信息")
	flag.Parse()

	if showVersion {
		fmt.Println("development-environment-cli 版本:", version)
		os.Exit(0)
	}
}
