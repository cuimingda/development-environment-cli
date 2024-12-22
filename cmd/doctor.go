package main

import (
	"fmt"
)

func HandleDoctorCommand(args []string) {
	// 输出参数
	fmt.Println("hello doctor")
	for _, arg := range args {
		fmt.Println(arg)
	}
}
