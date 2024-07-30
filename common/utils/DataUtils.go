package utils

import (
	"strings"
	"unicode"
)

/*
下划线转小驼峰
*/
func SnakeToLowerCamel(s string) string {
	components := strings.Split(s, "_")
	for i := 1; i < len(components); i++ {
		components[i] = strings.Title(components[i])
	}
	return strings.Join(components, "")
}

/*
下划线转大驼峰
*/
func SnakeToPascal(s string) string {
	words := strings.Split(s, "_")
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, "")
}

/*
判断字符串是否为空

	特别注意 空字符串和NULL字符串都是空字符串
*/
func IsEmpty(str string) bool {
	if len(str) == 0 {
		return true
	}
	upperCaseString := strings.ToUpper(str)
	if upperCaseString == "NULL" {
		return true
	}
	return false
}

/*
判断字符串是否不为空
*/
func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

/*
单词首字母小写
*/
func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
