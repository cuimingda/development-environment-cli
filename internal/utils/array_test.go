package utils

import (
	"testing"
)

func TestEnsureCommandArgsLength(t *testing.T) {
	tests := []TestCase{
		// 测试参数个数在范围内
		{EnsureCommandArgsLength, true, []interface{}{[]string{"arg1"}, 1}},
		{EnsureCommandArgsLength, true, []interface{}{[]string{}, 0}},

		{EnsureCommandArgsLength, false, []interface{}{[]string{"arg1", "arg2"}, 2}},
		{EnsureCommandArgsLength, false, []interface{}{[]string{"arg1", "arg2", "arg3"}, 3}},
		{EnsureCommandArgsLength, false, []interface{}{[]string{"arg1"}, 2}},
		{EnsureCommandArgsLength, false, []interface{}{[]string{"arg1", "arg2"}, 1}},
	}

	RunTests(t, tests)
}
