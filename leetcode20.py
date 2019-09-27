class Solution:
    def isValid(self, s: str) -> bool:
        hash = {'(':')', '{':'}', '[':']'}
        stack = []
        for char in s:
            char = str(char)
            if char == '(' or char == '{' or char == '[':
                stack.append(char)
            elif char == ')' or char == '}' or char == ']':
                if len(stack) == 0:
                    return False
                cpr = stack.pop()
                if char != hash.get(cpr):
                    return False
        if len(stack) != 0:
            return False
        return True

a = Solution()
a.isValid("[]{}")