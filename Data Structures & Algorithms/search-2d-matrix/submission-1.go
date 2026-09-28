func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		L, R := 0, len(row) - 1

		var mid int

		for L <= R {
			mid = (L + R) / 2
			if  target > row[mid]{
				L = mid + 1
			} else if  target < row[mid]{
				R = mid - 1
			} else {
				return true
			}
		}
	}
	return false 
}
