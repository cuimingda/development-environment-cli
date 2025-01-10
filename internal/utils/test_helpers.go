package utils

import (
	"bytes"
	"fmt"
	"path"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

// 定义一个通用的测试用例结构体
type TestCase struct {
	fn            interface{}
	expectSuccess bool
	args          []interface{}
}

type CommandTestCase struct {
	Args          []string
	ExpectSuccess bool
}

// 通用的测试辅助函数
func RunTests(t *testing.T, tests []TestCase) {
	for _, test := range tests {
		// 格式化测试名称为调用的参数列表
		argStrings := make([]string, len(test.args))

		for i, arg := range test.args {
			argStrings[i] = fmt.Sprintf("%v", arg)
		}

		testName := fmt.Sprintf(
			"%s(%s)=>%v",
			getFunctionName(test.fn),
			strings.Join(argStrings, ","),
			test.expectSuccess,
		)

		t.Run(testName, func(t *testing.T) {
			// 使用反射调用函数
			fn := reflect.ValueOf(test.fn)
			in := make([]reflect.Value, len(test.args))
			for i, arg := range test.args {
				in[i] = reflect.ValueOf(arg)
			}
			result := fn.Call(in)
			err, _ := result[0].Interface().(error)

			if test.expectSuccess {
				if err != nil {
					t.Errorf("%s returned error: %v", testName, err)
				}
			} else {
				if err == nil {
					t.Errorf("%s did not return error as expected", testName)
				}
			}
		})
	}
}

// 通用的命令测试辅助函数
func RunCommandTests(t *testing.T, rootCmd *cobra.Command, tests []CommandTestCase) {
	for _, test := range tests {
		// 重置输出缓冲区
		output := new(bytes.Buffer)
		rootCmd.SetOut(output)
		rootCmd.SetErr(output)

		// 设置命令行参数
		rootCmd.SetArgs(test.Args)

		// 执行命令
		err := rootCmd.Execute()

		// 检查错误
		if test.ExpectSuccess {
			if err != nil {
				t.Errorf("expected no error for args %v, but got %v", test.Args, err)
			}
		} else {
			if err == nil {
				t.Errorf("expected error for args %v, but got none", test.Args)
			}
		}
	}
}

func getFunctionName(i interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return path.Base(fullName)
}
