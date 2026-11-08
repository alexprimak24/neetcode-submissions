class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        unique = set()
        L = 0

        for R in range(len(nums)):
            if R - L > k:    
                unique.remove(nums[L])
                L = L + 1
            
            if nums[R] in unique:
                return True
            unique.add(nums[R])
        return False


    