func findMaxConsecutiveOnes(nums []int) int {
	maxConsecutiveOnes := 0
    count := 0
    for i := 0; i < len(nums); i++ {
        if (nums[i] == 1) {
            count += 1
        } else {
            maxConsecutiveOnes = max(maxConsecutiveOnes, count)
            count = 0;
        }
    }
    maxConsecutiveOnes = max(maxConsecutiveOnes, count)
    return maxConsecutiveOnes
}
