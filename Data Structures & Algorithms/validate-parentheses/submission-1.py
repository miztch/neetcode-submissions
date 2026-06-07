class Solution:
    def isValid(self, s: str) -> bool:
        pairs = {")": "(", "}": "{", "]": "["}

        stack = []
        for i,c in enumerate(s):
            if c in pairs.values():
                stack.append(c)
            else:
                if len(stack) == 0:
                    return False

                if stack[-1] == pairs[c]:
                    stack.pop(-1)
                else:
                    return False
        
        return len(stack) == 0