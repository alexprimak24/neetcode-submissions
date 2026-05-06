class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
      i = 0
      while i < len(nums):
         if nums[i] == val:
            for index in range(i + 1, len(nums)):
               nums[index - 1] = nums[index]
            i = 0
            nums.pop()
            continue
         i += 1
     
      return len(nums) 