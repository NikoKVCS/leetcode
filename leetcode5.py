class Solution:
    def longestPalindrome(self, s: str) -> str:
        s_new = "$#"
        for c in s:
            s_new += c
            s_new += "#"
        s_new += "&"
        n = len(s_new)
        mx = 0
        id = 0
        max_id = 0
        max_r = 1
        p = [0 for i in range(n)]
        for i in range(1,n-1):
            if i < mx:
                p[i] = min(p[2*id-i], mx-i)
            else:
                p[i] = 1
            while s_new[i-p[i]] == s_new[i+p[i]]:
                p[i] += 1
            if i+p[i] > mx:
                id = i
                mx = i+p[i]
            if p[i] >= max_r:
                max_id = i
                max_r = p[i]

        result = s_new[max_id-(max_r-1):max_id+(max_r-1)]
        result = result.replace("#","").replace("$","").replace("&","")
        return result

a = Solution()
print(a.longestPalindrome("abcba"))