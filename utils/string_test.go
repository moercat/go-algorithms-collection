package utils

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"racecar", "racecar"},
		{"hello世界", "界世olleh"}, // 测试多字节字符
	}

	for _, test := range tests {
		result := ReverseString(test.input)
		if result != test.expected {
			t.Errorf("ReverseString(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"aba", true},
		{"abba", true},
		{"hello", false},
		{"racecar", true},
		{"A man a plan a canal Panama", false}, // 因为有空格，所以不是回文
		{"raceacar", false},
	}

	for _, test := range tests {
		result := IsPalindrome(test.input)
		if result != test.expected {
			t.Errorf("IsPalindrome(%q) = %t; expected %t", test.input, result, test.expected)
		}
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"hello", 1},
		{"hello world", 2},
		{"  hello   world  ", 2}, // 测试多余的空格
		{"the quick brown fox jumps over the lazy dog", 9},
	}

	for _, test := range tests {
		result := CountWords(test.input)
		if result != test.expected {
			t.Errorf("CountWords(%q) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

func TestRemoveWhitespace(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"hello", "hello"},
		{"hello world", "helloworld"},
		{"  hello   world  ", "helloworld"},
		{"\thello\nworld\r", "helloworld"},
	}

	for _, test := range tests {
		result := RemoveWhitespace(test.input)
		if result != test.expected {
			t.Errorf("RemoveWhitespace(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestCapitalizeWords(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"hello", "Hello"},
		{"hello world", "Hello World"},
		{"HELLO WORLD", "Hello World"},
		{"hELLo WoRLd", "Hello World"},
		{"the quick brown fox", "The Quick Brown Fox"},
	}

	for _, test := range tests {
		result := CapitalizeWords(test.input)
		if result != test.expected {
			t.Errorf("CapitalizeWords(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestCommonPrefix(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{}, ""},
		{[]string{""}, ""},
		{[]string{"hello"}, "hello"},
		{[]string{"hello", "help", "hero"}, "he"},
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"interspecies", "interstellar", "interstate"}, "inters"},
	}

	for _, test := range tests {
		result := CommonPrefix(test.input)
		if result != test.expected {
			t.Errorf("CommonPrefix(%v) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{"", "", true},
		{"a", "a", true},
		{"ab", "ba", true},
		{"abc", "bca", true},
		{"listen", "silent", true},
		{"hello", "bello", false},
		{"hello", "world", false},
		{"abc", "def", false},
		{"aab", "abb", false},
	}

	for _, test := range tests {
		result := IsAnagram(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("IsAnagram(%q, %q) = %t; expected %t", test.s1, test.s2, result, test.expected)
		}
	}
}