// 所有fatal断言
package utils

func FatalIf(condition bool, format string, args ...any) {
	if condition {
		PrintErrorMessage(format, args...)
		ExitWithFailure()
	}
}

func FatalIfNot(condition bool, format string, args ...any) {
	if !condition {
		PrintErrorMessage(format, args...)
		ExitWithFailure()
	}
}

func FatalOnError(err error, format string, args ...any) {
	FatalIf(
		err != nil,
		format+" - %v",
		append(args, err)...,
	)
}
