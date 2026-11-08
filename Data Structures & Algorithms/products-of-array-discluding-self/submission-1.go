func productExceptSelf(nums []int) []int {
	preArr := make([]int, len(nums)+1)
	postArr := make([]int, len(nums)+1)
	preArr[0] = 1
	postArr[len(postArr) -1] = 1
	// calc prefix 
	for i := 0; i < len(nums); i++ {
		preArr[i+1] = preArr[i] * nums[i] 
	}
	
	// calc postfix
	for i := len(nums) -1; i >= 0; i-- {
		postArr[i] = postArr[i + 1] * nums[i]
	}

	productExSelf := make([]int, len(nums))
	for i, _ := range nums {
		productExSelf[i] = preArr[i] * postArr[i + 1]
	}

	return productExSelf
}
