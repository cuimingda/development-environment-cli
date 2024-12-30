// 基于语义对fmt进行二次封装
package utils

import (
	"fmt"
	"os"
)

func PrintInfoMessage(format string, args ...any) {
	fmt.Fprintf(os.Stdout, format, args...)
	fmt.Fprintln(os.Stdout)
}

func PrintSuccessMessage(format string, args ...any) {
	fmt.Fprintf(os.Stdout, format, args...)
	fmt.Fprintln(os.Stdout)
}

func PrintErrorMessage(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}
