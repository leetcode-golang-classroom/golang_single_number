package sol

func singleNumber(nums []int) int {
	result := nums[0]
	lenNums := len(nums)
	for pos := 1; pos < lenNums; pos++ {
		result ^= nums[pos]
	}
	return result
}
