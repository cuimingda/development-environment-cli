package utils

import "testing"

func TestEnsureIntegerRange(t *testing.T) {
	tests := []TestCase{
		{EnsureIntegerRange, true, []interface{}{5, 1, 10}},
		{EnsureIntegerRange, true, []interface{}{1, 1, 10}},
		{EnsureIntegerRange, true, []interface{}{10, 1, 10}},
		{EnsureIntegerRange, false, []interface{}{11, 1, 10}},
		{EnsureIntegerRange, false, []interface{}{0, 1, 10}},
		{EnsureIntegerRange, false, []interface{}{-1, 1, 10}},
		{EnsureIntegerRange, false, []interface{}{5, 10, 1}},
	}

	RunTests(t, tests)
}
