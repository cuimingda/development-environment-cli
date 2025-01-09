package utils

import "testing"

func TestEnsureIntegerRange(t *testing.T) {
	tests := []TestCase{
		{EnsureIntegerRange, true, []interface{}{5, 1, 10}},
		{EnsureIntegerRange, false, []interface{}{11, 1, 10}},
	}

	RunTests(t, tests)
}
