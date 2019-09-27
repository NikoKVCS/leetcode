class Solution:
    def singleNumber(self, nums) -> int:
        map = dict()
        for item in nums:
            if item in map.keys():
                map.pop(item)
            else:
                map[item] = item
        return map.popitem()[0]
a = Solution()
print(a.singleNumber([4,1,2,1,2]))