class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        count = {}

        occur = [[] for i in range(len(nums) + 1)]

        for num in nums:
            count[num] = 1 + count.get(num, 0)

        for key, val in count.items():
            occur[val].append(key)
        
        mostFrequest = []
        for numOccur in range(len(occur) - 1, -1, -1):
            for num in occur[numOccur]:
                mostFrequest.append(num)
                # we check len here, as can happen that we never
                # return in case appended enough only thanks to
                # the last element
                if len(mostFrequest) == k:
                    return mostFrequest
        


        