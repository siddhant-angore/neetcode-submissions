func calPoints(operations []string) int {
	// begin with empty record
	// score := 0
	// let's make a board to keep track of records
	record := []int{}
	for _, op := range operations {
		if op == "+" {
			// +: record new score i.e. sum of previous 2 scores
			record = append(record, record[len(record)-1] + record[len(record)-2])
		} else if op == "D" {
			// D: record new score i.e. double (2x) of previous score
			record = append(record, record[len(record)-1] * 2)
		} else if op == "C" {
			// C: remove previous score from record
			record = record[:len(record)-1]
		} else {
			// x: record new score x
			num, err := strconv.Atoi(op)
			if err != nil {
				fmt.Println("error")
				return 0
			}
			record = append(record, num)
		}				
	}
	return addRecord(record)
}

func addRecord(record []int) int {
	sum := 0
	for _, r := range record {
		sum += r
	}
	return sum
}
