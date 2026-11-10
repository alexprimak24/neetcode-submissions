class Solution:
    def pivotIndex(self, nums: List[int]) -> int:
        n = len(nums)
        right = [0] * n
        left = [0] * n

        totalRight = 0
        # calc right + 1
        for i in range(n):
            totalRight = totalRight + nums[i]
            right[i] = totalRight
            

        # calc left
        totalLeft = 0
        for i in range(n - 1, -1, -1 ):
            totalLeft = totalLeft + nums[i]
            left[i] = totalLeft

        # look matches
        for i in range(n):
            if right[i] == left[i]:
                return i

        # not found -1
        return -1

        