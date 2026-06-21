func calPoints(operations []string) int {
	ops := []int{}
	sum := 0
	for _, op := range operations {
		if op == "+" {
			newVal := ops[len(ops) -1] + ops[len(ops)-2]
			ops = append(ops, newVal)
			sum += newVal
			continue
		}
		if op == "D" {
			newVal := ops[len(ops) -1] * 2
			ops = append(ops,newVal)
			sum += newVal
			continue
		}
		if op == "C" {
			lastVal := ops[len(ops)-1]
			ops = ops[:len(ops)-1]
			sum -= lastVal
	 		continue
		}
		val, err := strconv.Atoi(op)
		if err == nil {
			ops = append(ops, val)
			sum += val
			continue
		}
	}

	return sum
}
