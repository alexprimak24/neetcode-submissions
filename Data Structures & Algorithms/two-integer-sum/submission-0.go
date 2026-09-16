func twoSum(nums []int, target int) []int {
    needed := map[int]int{}

	returnedArr := make([]int,0)
	for i, num := range nums {
		if val, ok := needed[num]; ok {
			returnedArr = append(returnedArr, val)
			returnedArr = append(returnedArr, i)
			return returnedArr
		}
		needed[target-num] = i
	}
	return returnedArr
}
