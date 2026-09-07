func calPoints(operations []string) int {
	res := []int{}

	for _, operation := range operations {
		operationInt, err := strconv.Atoi(operation)
		if err == nil {
			res = append(res,operationInt)
			continue
		}

		if operation == "D" {
			last := res[len(res)-1]

			res = append(res, last * 2)
			continue
		}

		if operation == "+" {
			last := res[len(res)-1]
			prevLast := res[len(res)-2]

			res = append(res, last+prevLast)
			continue
		}

		if operation == "C" {
			res = res[:len(res)-1]
			continue
		}
	}

	sum := 0
	for _, v := range res {
		sum += v
	}
	return sum
}
