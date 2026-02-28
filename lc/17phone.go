package lc

var m = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {

	var res []string
	for i := 0; i < len(digits); i++ {
		res = AppendStr(m[string(digits[i])], res)

	}
	return res
}

func AppendStr(enum, res []string) []string {
	if len(res) == 0 {
		return enum
	}
	var newRes []string
	for i := 0; i < len(res); i++ {
		for _, v := range enum {
			newRes = append(newRes, res[i]+v)
		}
	}

	return newRes
}
