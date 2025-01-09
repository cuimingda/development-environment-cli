package utils

import "fmt"

func EnsureIntegerRange(number int, min int, max int) error {
	if min > max {
		return fmt.Errorf("最小值 %d 大于最大值 %d", min, max)
	}
	if number < min || number > max {
		return fmt.Errorf("数值超出范围：期望 %d-%d，实际 %d", min, max, number)
	}
	return nil
}
