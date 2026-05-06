class Solution:
    def replaceElements(self, arr: List[int]) -> List[int]:

        i = 0
        n = len(arr)

        biggest = 0
        while i < n:
            for index in range(i + 1, n):
                if arr[index] > biggest:
                    biggest = arr[index]
            arr[i] = biggest
            biggest = 0
            i += 1

        arr[n-1] = -1

        return arr

        