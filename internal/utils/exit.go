// 语义化封装os.exit
package utils

import "os"

func ExitWithFailure() {
	os.Exit(1)
}

func ExitWithSuccess() {
	os.Exit(0)
}
