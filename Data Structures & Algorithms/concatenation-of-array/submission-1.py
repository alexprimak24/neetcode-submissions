class Solution:
    def getConcatenation(self, nums: List[int]) -> List[int]:
        length = len(nums)

        newArr = [0] * length * 2

        for i in range(length):
            newArr[i] = newArr[i + length] = nums[i]

        return newArr