class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
       pairs = {}

       for i in range(len(nums)):
        num = nums[i]
        if num in pairs:
            return [pairs[num], i]
        else:
            pairs[target-num] = i
