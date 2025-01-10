package utils

import "fmt"

// 参数长度范围常量
const (
	MinArgsLength = 0
	MaxArgsLength = 1
)

func EnsureCommandArgsLength(args []string, length int) error {

	err := EnsureIntegerRange(length, MinArgsLength, MaxArgsLength)
	if err != nil {
		return fmt.Errorf("参数个数范围不正确：%v", err)
	}

	if len(args) != length {
		return fmt.Errorf("参数个数不正确，期望 %d 个参数，实际 %d 个参数", length, len(args))
	}

	return nil
}
