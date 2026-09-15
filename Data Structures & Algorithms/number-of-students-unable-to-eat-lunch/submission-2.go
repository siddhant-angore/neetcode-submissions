func countStudents(students []int, sandwiches []int) int {
    res := len(students)
	cnt := make([]int, 2)
	for _, student := range students {
		cnt[student]++
	}
	for _, sandwich := range sandwiches {
		if cnt[sandwich] > 0 {
			cnt[sandwich]--
			res--
		} else {
			break
		}
	}
	return res
}