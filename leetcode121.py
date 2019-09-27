class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        n = len(prices)
        best = 0
        for i in range(n-1, -1, -1):
            if i == 0:
                break
            for j in range(i-1, -1, -1):
                if prices[i] - prices[j] > best:
                    best = prices[i] - prices[j]
        return best
a = Solution()
print(a.maxProfit([7,1,5,3,6,4]))
print(a.maxProfit([7,6,4,3,1]))