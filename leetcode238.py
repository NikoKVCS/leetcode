class Solution:
    def productExceptSelf(self, nums):
        if len(nums) <= 1:
            return nums
        dp = [1 for i in range(0,len(nums))]
        result = [1 for i in range(0,len(nums))]
        dp[0] = 1
        result[len(nums)-1] = 1
        for i in range(len(nums)-2, -1, -1):
            result[i] = result[i+1] * nums[i+1]
        for i in range(1,len(nums)):
            dp[i] = dp[i-1] * nums[i-1]
            result[i] *= dp[i]
        result[0] *= dp[0]
        return result

a = Solution()
print(a.productExceptSelf([1,2,3,4]))