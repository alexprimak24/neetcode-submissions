func search(nums []int, target int) int {
	L, R := 0, len(nums) - 1

	var mid int
	for L <= R {
		mid = (L + R) / 2

		if target > nums[mid] {
			L = mid + 1
		} else if target < nums[mid] {
			R = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
