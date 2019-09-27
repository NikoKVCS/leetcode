class Solution:
    def subsets(self, nums) :
        result = [[]]
        if len(nums) == 0:
            return result
        for i in range(1, len(nums)+1):
            lists = self.C(nums, i)
            for item in lists:
                result.append(item)
        return result
    
    def C(self, nums, n):
        result = []
        if n == 1:
            for item in nums:
                result.append([item])
            return result
        while len(nums) >= n:
            nums = nums.copy()
            a = nums.pop()
            lists = self.C(nums,n-1)
            for item in lists:
                item.append(a)
                result.append(item)
        return result
                
a = Solution()
print(a.subsets([1,2,3,4]))
