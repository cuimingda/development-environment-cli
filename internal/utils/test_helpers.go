package utils

import (
	"fmt"
	"path"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

// 定义一个通用的测试用例结构体
type TestCase struct {
	fn            interface{}
	expectSuccess bool
	args          []interface{}
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

func getFunctionName(i interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return path.Base(fullName)
}
