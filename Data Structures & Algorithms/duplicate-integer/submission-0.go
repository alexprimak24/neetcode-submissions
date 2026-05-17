func hasDuplicate(nums []int) bool {
    unique := make(map[int]bool)

    for _, num := range nums {
        _, ok := unique[num]
        if ok {
            return true 
        } else {
            unique[num] = true
        }
    }

    return false
}
