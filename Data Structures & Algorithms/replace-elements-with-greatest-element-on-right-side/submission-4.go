func replaceElements(arr []int) []int {
	currMax := -1

	// loop in reverse order
	for i := len(arr)-1; i >= 0; i-- {
		newMax := arr[i]
		if newMax < currMax {
			newMax = currMax
		}
		

		arr[i] = currMax

		// new max is compare old max with prev value
		currMax = newMax
	}

	return arr

}
