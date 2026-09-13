func getConcatenation(nums []int) []int {
	length := len(nums)
    ans := make([]int, 2 * length)
	for i, n := range nums {
		ans[i], ans[i + length] = n, n
	}
	return ans
}
