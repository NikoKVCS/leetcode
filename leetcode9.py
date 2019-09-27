class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0 :
            return False
        t = x
        y = 0
        while t // 10  or t % 10:
            y = y * 10 + t%10
            t = t // 10
        if x == y:
            return True
        return False