func findMaxConsecutiveOnes(nums []int) int {
    max := 0
    curr := 0

    for _, num := range nums {
       if num == 1 {
        curr++
       } else {
        curr = 0
       }
       if curr > max {
        max = curr
    }
    }

    
    return max
}
