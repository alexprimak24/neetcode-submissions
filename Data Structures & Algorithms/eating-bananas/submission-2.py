class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        # find max
        maxi = max(piles)
        # L and R = 0, max
        L, R = 1, maxi
        # assign min  
        res = maxi    
        # start from binary search
        # iter until found
        while L <= R:
            mid = (L + R) // 2

            hours = 0
            for pile in piles:
                hours += math.ceil(pile / mid) 
           
            if hours <= h:
                res = mid
                R = mid - 1
            else:
                L = mid + 1
        return res