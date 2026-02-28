package lc41

func firstMissingPositive(nums []int) int {

	var cm = make(map[int]struct{})
	var tmpNum = 1
	for i := 0; i < len(nums); i++ {
		cm[nums[i]] = struct{}{}
	}
	for {
		if _, ok := cm[tmpNum]; !ok {
			return tmpNum
		}
		tmpNum++
	}
}
