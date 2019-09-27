class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        if m < 1 or n < 1:
            return 0
        sum_up = 1
        for i in range(1,m+n-1):
            sum_up *= i
        sum_m = 1
        for i in range(1,m):
            sum_m *= i
        sum_n = 1
        for i in range(1,n):
            sum_n *= i
        return int(sum_up / (sum_m*sum_n))
a = Solution()
print(a.uniquePaths(3,2))