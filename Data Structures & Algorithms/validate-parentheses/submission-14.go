func isValid(s string) bool {
	st := []string{}
	m := map[string]string{"(": ")", "[": "]", "{": "}"}
	for _, c := range s {
		ch := string(c)
		fmt.Println(ch)
		if _, ok := m[ch]; ok {
			st = append(st, m[ch])
		} else {
			if len(st) > 0 && ch == st[len(st)-1] {
				st = st[:len(st)-1]
			} else {
				return false
			}
		}
	}
	return len(st) == 0
}
