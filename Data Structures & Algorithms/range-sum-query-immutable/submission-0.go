type NumArray struct {
	prefixSums []int
}

func Constructor(nums []int) NumArray {
	prefixSums := make([]int, len(nums))
	prefixSums[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		prefixSums[i] = nums[i] + prefixSums[i-1]
	}

	return NumArray{
		prefixSums: prefixSums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	leftSum := 0
	if left > 0 {
		leftSum = this.prefixSums[left-1]
	}
	return this.prefixSums[right] - leftSum
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */