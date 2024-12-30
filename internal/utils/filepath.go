// 对filepath的封装
package utils

import (
	"os"
	"path/filepath"
)

func GetWorkingDirOrFatal() string {
	workingDir, err := os.Getwd()
	FatalOnError(err, "获取当前工作目录失败")

	return workingDir
}

func GetWorkingDirName() string {
	workingDir := GetWorkingDirOrFatal()
	dirName := filepath.Base(workingDir)

	return dirName
}

func CreateDir(dir string) {
	PrintActionLog("正在创建目录 %s...", dir)
	err := os.MkdirAll(dir, os.ModePerm)
	FatalOnError(err, "创建目录 %s 失败", dir)
	PrintActionLog("[INFO] 目录 %s 创建成功\n", dir)
}

func EnterDir(dir string) {
	PrintActionLog("正在切换到目录 %s...", dir)
	err := os.Chdir(dir)
	FatalOnError(err, "切换到目录 %s 失败", dir)
	PrintActionLog("切换到目录 %s 成功", dir)
}
