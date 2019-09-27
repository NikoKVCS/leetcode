class Solution:
    def myAtoi(self, str: str) -> int:
        flag = 1
        result = 0
        id = 0
        str = str.strip()
        for a in str:
            if ord(a) >= ord('0') and ord(a) <= ord('9'):
                result = result * 10 + ord(a) - ord('0')
            elif a == '-' and 0 == id:
                flag = -1*flag
            elif a == '+' and 0 == id:
                pass
            else:
                break
            id += 1
        result = min(result*flag,2**31-1) if result * flag > 0 else max(result*flag,-2**31) 
        return result

a = Solution()
print(a.myAtoi("   -42"))