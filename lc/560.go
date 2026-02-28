package lc

func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	m2 := map[int]int{}
	m2[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m2[pre-k]; ok {
			count += m2[pre-k]
		}
		m2[pre] += 1
	}
	return count
}
