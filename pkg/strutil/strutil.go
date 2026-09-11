// Package strutil 提供字符串处理工具函数，作为 pkg/ 公共工具包的示例。
package strutil

import "strings"

// Reverse 按 Unicode 字符反转字符串。
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsBlank 判断字符串为空或仅包含空白字符。
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}
