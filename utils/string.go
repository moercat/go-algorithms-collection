package utils

import (
	"strings"
	"unicode"
)

// ReverseString 反转字符串
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome 检查字符串是否是回文
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	
	return true
}

// CountWords 计算字符串中的单词数
func CountWords(s string) int {
	fields := strings.Fields(s)
	return len(fields)
}

// RemoveWhitespace 移除字符串中的所有空白字符
func RemoveWhitespace(s string) string {
	var result strings.Builder
	for _, r := range s {
		if !unicode.IsSpace(r) {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// CapitalizeWords 将每个单词的首字母大写
func CapitalizeWords(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, " ")
}

// CommonPrefix 求字符串数组的最长公共前缀
func CommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	
	prefix := strs[0]
	for _, s := range strs[1:] {
		for !strings.HasPrefix(s, prefix) {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	
	return prefix
}

// IsAnagram 检查两个字符串是否是字母异位词
func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	
	count := make(map[rune]int)
	for _, r := range s1 {
		count[r]++
	}
	
	for _, r := range s2 {
		count[r]--
		if count[r] < 0 {
			return false
		}
	}
	
	return true
}