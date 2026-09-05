func replaceElements(arr []int) []int {
	for i, _ := range arr {
		if i + 1 == len(arr){
			arr[i] = -1
			return arr
		}


		greatest := 0
		
		for j:= i+1; j < len(arr); j++  {
			if arr[j] > greatest {
				greatest = arr[j]
			}
		}

		

		arr[i] = greatest
		greatest = 0
	}

	return arr
}
