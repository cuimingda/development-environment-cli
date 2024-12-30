// 字符串处理函数
package utils

import "strings"

func TrimmedStringFromBytes(b []byte) string {
	return strings.TrimSpace(string(b))
}
