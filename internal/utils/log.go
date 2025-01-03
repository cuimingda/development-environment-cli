// 基于语义对log进行二次封装
// 认为规则：所有的log都受verbose控制
package utils

import "log"

var Verbose bool

func PrintActionLog(format string, args ...any) {
	if Verbose {
		log.Printf(format, args...)
	}
}

func PrintInfoLog(format string, args ...any) {
	if Verbose {
		log.Printf(format, args...)
	}
}

func PrintSuccessLog(format string, args ...any) {
	if Verbose {
		log.Printf(format, args...)
	}
}

func PrintErrorLog(format string, args ...any) {
	if Verbose {
		log.Printf(format, args...)
	}
}
