package main

import (
	"fmt"
	"os"
)

func main() {

	// 检查子命令
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "doctor":
			HandleDoctorCommand()
			return
		}
	}

	fmt.Println("hello world")
}
