func findMaxConsecutiveOnes(nums []int) int {
    max := 0
    curr := 0
	for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            curr = 0
        } else {
            curr++
            if curr > max {
                max = curr
            }
        }
    }
    return max
}
