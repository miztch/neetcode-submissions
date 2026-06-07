
class Solution:
    def isValid(self, s: str) -> bool:
        openings = ["(","{","["]
        closings = [")","}","]"]

        stack = []
        for i,c in enumerate(s):
            if c in openings:
                stack.append(c)
            else:
                if len(stack) == 0:
                    return False

                if stack[-1] == openings[closings.index(c)]:
                    stack.pop(-1)
                else:
                    return False
        
        return len(stack) == 0