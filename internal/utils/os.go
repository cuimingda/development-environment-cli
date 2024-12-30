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
