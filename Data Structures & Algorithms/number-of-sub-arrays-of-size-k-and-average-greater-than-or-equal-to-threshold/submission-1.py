class Solution:
    def numOfSubarrays(self, arr: List[int], k: int, threshold: int) -> int:
        threshold = threshold * k

        acc = res = 0

        for L in range(len(arr)):
            acc += arr[L]

            if L >= k - 1:
                if acc >= threshold:
                    res += 1
                acc -= arr[L - k + 1]
        return res

            