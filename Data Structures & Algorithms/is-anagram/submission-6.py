class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        freq = {}

        for char1 in s:
            freq[char1] = freq.get(char1, 0) + 1

        for char2 in t:
            freq[char2] = freq.get(char2, 0) - 1

        for value in freq.values():
            if value != 0:
                return False
        return True
