package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

func debugOsArgs() {
	log.Printf("typeof of os.Args - %s\n", reflect.TypeOf(os.Args))

	log.Printf("len(os.Args) - %d", len(os.Args))
	for i, arg := range os.Args {
		log.Printf("os.Args[%d] - %s", i, arg)
	}
}

func main() {

	debugOsArgs()

	// 检查子命令
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "doctor":
			HandleDoctorCommand()
			return
		case "web":
			HandleWebCommand(os.Args[2:])
			return
		}
	}

	fmt.Println("hello world")
}
