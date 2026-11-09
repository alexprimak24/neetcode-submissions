class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        maxProf = 0
        L = 0

        for R in range(L + 1,len(prices)):
            if prices[R] < prices[L]:
                L = R
            else:
                if maxProf < prices[R] - prices[L]:
                    maxProf = prices[R] - prices[L]
                



        return  maxProf
