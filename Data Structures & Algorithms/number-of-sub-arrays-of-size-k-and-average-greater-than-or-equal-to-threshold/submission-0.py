class Solution:
    def numOfSubarrays(self, arr: List[int], k: int, threshold: int) -> int:
        threshold *= k
        res = currSum = 0

        for R in range(len(arr)):
            currSum += arr[R]
            if R >= k - 1:
                if currSum >= threshold:
                    res += 1
                currSum -= arr[R - k + 1]
        return res