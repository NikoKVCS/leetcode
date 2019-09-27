class Solution:
    def climbStairs(self, n: int) -> int:
        max_d = n // 2
        total = 0
        for i in range(0,max_d+1):
            M = i + n - 2*i
            N = M - i if M - i > i else i
            sum_up = 1
            for j in range(N+1, M+1):
                sum_up *= j
            sum_do = 1
            for j in range(1, M-N+1):
                sum_do *= j
            total += int(sum_up/sum_do)
            print(total)
        return total

a = Solution()
a.climbStairs(44)