func findMaxConsecutiveOnes(nums []int) int {
    max := 0
    curr := 0

    for _, num := range nums {
        if num == 0 {
            if curr > max {
                max = curr
            }
            curr = 0
            continue
        }
        curr++
    }

    if curr > max {
        max = curr
    }
    return max
}
