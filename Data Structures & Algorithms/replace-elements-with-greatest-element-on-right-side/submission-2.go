func replaceElements(arr []int) []int {
	maxVal := arr[len(arr) - 1]
	arr[len(arr) - 1] = -1
	for i := len(arr) - 2; i > -1; i-- {
		val := arr[i]
		arr[i] = maxVal
		maxVal = max(maxVal, val)
	}
	return arr
}
