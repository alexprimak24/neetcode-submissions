func calPoints(operations []string) int {
	ops := []int{}
	for _, op := range operations {
		

		if op == "+" {
			ops = append(ops, ops[len(ops) -1] + ops[len(ops)-2] )
			continue
		}
		if op == "D" {
			ops = append(ops,ops[len(ops) -1] * 2)
			continue
		}
		if op == "C" {
			ops = ops[:len(ops)-1]
	 		continue
		}
		val, err := strconv.Atoi(op)
		if err == nil {
			ops = append(ops, val)
			continue
		}
	}
	sum := 0
	for _, op := range ops {
		sum += op
	}

	return sum
}
