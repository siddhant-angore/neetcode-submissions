func removeElement(nums []int, val int) int {
	count := 0
	length := len(nums)
	i := 0
    for i < length {
		if nums[i] == val {
			removeIndex(nums[:], i, length - count)
			count += 1	
		} else {
			i += 1
		}
	}
	return length - count
}

func removeIndex(nums []int, deletionIndex, length int) {
	// get the removal index
	// shift all the elements from index to end 1 position left
	for i := deletionIndex + 1; i < length; i++ {
		nums[i - 1] = nums[i]
	}
	// mark last element as 0
	nums[length - 1] = -1
}
