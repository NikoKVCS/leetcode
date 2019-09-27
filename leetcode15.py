class Solution:
    def threeSum(self, nums):
        result = []
        for i in range(0,len(nums)-2):
            map = dict()
            for j in range(i+1,len(nums)):
                if map.get(nums[j]) != None:
                    map.get(nums[j]).sort()
                    result.append(map.get(nums[j]))
                    map[nums[j]] = None
                else:
                    c = 0-nums[i]-nums[j]
                    map[c] = list([nums[i],nums[j],c])
        result_sort = []
        for item in result:
            append = True
            for i in range(0,len(result_sort)):
                if item == result_sort[i]:
                    append = False
                    break
                if item[0] < result_sort[i][0] or (item[0] == result_sort[i][0] and item[1] < result_sort[i][1]) or (item[0] == result_sort[i][0] and item[1] == result_sort[i][1] and item[2] < result_sort[i][2]):
                    result_sort.insert(i,item)
                    append = False
                    break
            if append:
                result_sort.append(item)
        return result_sort

a = Solution()
print(a.threeSum([0,2,2,3,0,1,2,3,-1,-4,2]))