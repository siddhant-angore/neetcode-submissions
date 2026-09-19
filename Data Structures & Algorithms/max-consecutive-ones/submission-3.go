func findMaxConsecutiveOnes(nums []int) int {
	maxOnes := 0
    count := 0
    for i := 0; i < len(nums); i++ {        
        if nums[i] == 1 {
            count++
        } else {
            if count > maxOnes {
                maxOnes = count
            }
            count = 0
        }
    }
    if count > maxOnes {
        maxOnes = count
    }
    return maxOnes
}
