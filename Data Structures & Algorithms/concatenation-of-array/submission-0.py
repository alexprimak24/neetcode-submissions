class Solution:
    def getConcatenation(self, nums: List[int]) -> List[int]:
        new_len = len(nums) * 2

        newArr = [0] * new_len

        for i in range(new_len):
            newArr[i] = nums[i%len(nums)]

        return newArr