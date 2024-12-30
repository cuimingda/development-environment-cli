package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ConfirmYesOrNo(format string, args ...any) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(format+"(Y/n)", args...)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	yes := input == "" || strings.ToLower(input) == "yes" || strings.ToLower(input) == "y"
	return yes
}
