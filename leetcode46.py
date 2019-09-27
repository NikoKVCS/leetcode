class Solution:
    def permute(self, nums:list):
        if len(nums) == 1:
            return [nums]
        all_list = []
        for a in nums:
            num_cpy = nums.copy()
            num_cpy.remove(a)
            nlist = self.permute(num_cpy)
            for l in nlist:
                l.append(a)
                all_list.append(l)
        return all_list

a = Solution()
print(a.permute([1,2,3]))
