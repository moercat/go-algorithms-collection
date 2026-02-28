package algorithm

import "fmt"

var fMap = map[uint8]bool{
	'+': true,
	'-': true,
	'*': true,
	'/': true,
}

var numMap = map[uint8]bool{
	'0': true,
	'1': true,
	'2': true,
	'3': true,
	'4': true,
	'5': true,
	'6': true,
	'7': true,
	'8': true,
	'9': true,
}

func beforeCal(s string) float64 {
	var numList []float64
	for i := len(s) - 1; i >= 0; i-- {
		if numMap[s[i]] {
			numList = append([]float64{float64(s[i] - '0')}, numList...)
		}
		if fMap[s[i]] {
			// 弹出最后两个数字
			num1 := numList[0]
			num2 := numList[1]
			var tmpNum float64
			if s[i] == '+' {
				tmpNum = num1 + num2
			}
			if s[i] == '-' {
				tmpNum = num1 - num2
			}
			if s[i] == '*' {
				tmpNum = num1 * num2
			}
			if s[i] == '/' {
				tmpNum = num1 / num2
			}
			numList = append([]float64{tmpNum}, numList[2:]...)
		}
	}
	return numList[0]
}

func front(s string) string {
	ft, num := []uint8{}, []uint8{}
	ftNums := map[uint8]int{
		'*': 2,
		'/': 2,
		'+': 1,
		'-': 1,
	}

	for i := len(s) - 1; i >= 0; i-- {
		if numMap[s[i]] {
			num = append(num, s[i])
			continue
		}
		if s[i] == '(' {
			for len(ft) >= 1 && ft[len(ft)-1] != ')' {
				num = append(num, ft[len(ft)-1])
				ft = ft[:len(ft)-1]
			}
			ft = ft[:len(ft)-1]
			continue
		}
		if s[i] == ')' {
			ft = append(ft, s[i])
			continue
		}
		a := string(s[i])
		if fMap[s[i]] {
			if len(ft) >= 1 && ft[len(ft)-1] == ')' {
				ft = append(ft, s[i])
				continue
			}
			if len(ft) == 0 || ftNums[s[i]] >= ftNums[ft[len(ft)-1]] {
				ft = append(ft, s[i])
				continue
			}
			for len(ft) >= 1 && ftNums[s[i]] < ftNums[ft[len(ft)-1]] {
				num = append(num, ft[len(ft)-1])
				ft = ft[:len(ft)-1]
			}
			ft = append(ft, s[i])
		}
		fmt.Println(a)
	}
	for len(ft) >= 1 && ft[len(ft)-1] != ')' {
		num = append(num, ft[len(ft)-1])
		ft = ft[:len(ft)-1]
	}

	var res string
	for i := len(num) - 1; i >= 0; i-- {
		res += string(num[i])
	}

	return res
}
