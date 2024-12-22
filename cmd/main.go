package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    // 定义子命令
    doctorCmd := flag.NewFlagSet("doctor", flag.ExitOnError)

    // 检查子命令
    if len(os.Args) > 1 {
        switch os.Args[1] {
        case "doctor":
            doctorCmd.Parse(os.Args[2:])
            HandleDoctorCommand(doctorCmd.Args())
            return
        }
    }

    fmt.Println("hello world")
}