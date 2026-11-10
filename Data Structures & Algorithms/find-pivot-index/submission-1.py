class Solution:
    def pivotIndex(self, nums: List[int]) -> int:
        rightSum = sum(nums)

        leftSum = 0
        for i in range(len(nums)):
            if rightSum - nums[i] - leftSum == leftSum:
                return i
            leftSum = leftSum + nums[i]
        return -1

        