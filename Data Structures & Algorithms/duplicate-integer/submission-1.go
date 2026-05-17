func hasDuplicate(nums []int) bool {
    unique := make(map[int]struct{})

    for _, num := range nums {
        _, ok := unique[num]
        if ok {
            return true 
        } else {
            unique[num] = struct{}{}
        }
    }

    return false
}
