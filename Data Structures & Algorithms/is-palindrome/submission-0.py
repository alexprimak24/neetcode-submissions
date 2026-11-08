class Solution:
    def isPalindrome(self, s: str) -> bool:
        L, R = 0, len(s) - 1
        while L < R:
            if not self.isAlphanumeric(s[L]):
                L += 1
                continue
            if not self.isAlphanumeric(s[R]):
                R -= 1
                continue
            if s[L].lower() != s[R].lower():
                return False
            L += 1
            R -= 1
        return True
    def isAlphanumeric(self, ch: str) -> bool:
        return ch.isalnum()

