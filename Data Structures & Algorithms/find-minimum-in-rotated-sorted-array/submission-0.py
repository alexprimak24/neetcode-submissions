class Solution:
    def findMin(self, nums: List[int]) -> int:
        res = nums[0]

        L, R = 0, len(nums) - 1

        while L <= R:
            if nums[L] < nums[R]:
                res = min(res, nums[L])
                break

            K = (L + R) // 2
            res = min(res, nums[K])
            if nums[K] >= nums[L]:
                L = K + 1
            else:
                R = K - 1

        return res