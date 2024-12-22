package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "1.0.0"

func main() {
	var showVersion bool
	var shortVersion bool

	flag.BoolVar(&showVersion, "version", false, "显示版本信息")
	flag.BoolVar(&shortVersion, "v", false, "显示版本信息")
	flag.Parse()

	if showVersion || shortVersion {
		fmt.Println("development-environment-cli 版本:", version)
		os.Exit(0)
	}

	fmt.Println("hello world again")
}
