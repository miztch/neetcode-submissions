class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        S = "".join(sorted([s[i] for i in range(len(s))]))
        T = "".join(sorted([t[i] for i in range(len(t))]))

        return S == T
