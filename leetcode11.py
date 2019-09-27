class Solution:
    def maxArea(self, height: List[int]) -> int:
        i = 0
        j = len(height) - 1
        max_a = 0
        while i != j :
            h = height[i] if height[i] < height[j] else height[j]
            total = h * (j-i)
            max_a = total if total > max_a else max_a
            if height[i] > height[j] :
                j -= 1
            else:
                i += 1
        return max_a