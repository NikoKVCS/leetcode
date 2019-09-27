class Solution:

    def reverse(self, x: int) -> int:
        t = x
        y = 0
        while 0!=t//10 or 0!=t%10:
            y = y*10 + t%10
            t = t//10
        if 2**31-1 < y or 0-2**31>y:
            y = 0
        return y

a = Solution()
print(a.reverse(123456))