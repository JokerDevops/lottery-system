package comm

import (
	"fmt"
	"testing"
)

func TestStripsSlashs(t *testing.T) {
	// 定义测试用例
	testCases := []struct {
		input    string
		expected string
	}{
		{"test\\", "test"},
		// 添加更多测试用例...
	}

	// 遍历测试用例
	for _, tc := range testCases {
		// 调用被测试的函数
		result := StripsSlashs(tc.input)

		// 检查结果是否符合预期
		if result != tc.expected {
			t.Errorf("StripsSlashs(%s) 返回 %s，期望 %s", tc.input, result, tc.expected)
		}
	}
}

func TestIp4toInt(t *testing.T) {
	s := Ip4toInt("127.0.0.1")
	fmt.Println(s)

}
