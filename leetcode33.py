class Solution:
    def search(self, nums, target: int) -> int:
        n = len(nums)
        if n <= 0:
            return 0
        left = 0
        right = n-1
        mid = (left + right) // 2
        if target >= nums[0]:#在左边
            while right - left > 1:
                if nums[mid] == target:
                    return mid
                elif nums[mid] < nums[0]:#在右区
                    right_temp = mid 
                    mid_temp = (right_temp + left) // 2
                    if nums[mid_temp] >= nums[0] and nums[mid_temp] < target:#在左区 且在target左边
                        left = mid 
                        mid = (left + right) // 2
                    else:
                        right = right_temp
                        mid = mid_temp
        else:
            while right - left > 1:
                if nums[mid] == target:
                    return mid
                elif nums[mid] >= nums[0]:#在左区
                    left_temp = mid 
                    mid_temp = (left_temp + right) // 2
                    if nums[mid_temp] < nums[0] and nums[mid_temp] > target:#在右区 且在target左边
                        right = mid 
                        mid = (left + right) // 2
                    else:
                        left = left_temp
                        mid = mid_temp

        if nums[left] == target:
            return left
        elif nums[right] == target:
            return right
        else:
            return -1
            
            
        
