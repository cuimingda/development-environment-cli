// 所有断言类检测，如果不符合条件则直接退出程序
package utils

func ensure(condition bool, format string, args ...any) {
	FatalIfNot(condition, format, args...)
}

func EnsureCommand(command string) {
	ensure(
		IsCommandAvailable(command),
		"命令 %s 未安装",
		command,
	)
}

func EnsurePath(path string) {
	ensure(
		IsPathAvailable(path),
		"文件或目录 %s 不存在",
		path,
	)
}

func EnsureAlpine() {
	ensure(
		IsAlpine(),
		"当前系统不是 %s",
		"Alpine Linux",
	)
}

func EnsureMacOS() {
	ensure(
		IsMacOS(),
		"当前系统不是 %s",
		"MacOS",
	)
}

func EnsureGitRemoteOrigin(url string) {
	ensure(
		IsGitRemoteOrigin(url),
		"当前git项目的remote origin不是 %s",
		url,
	)
}
