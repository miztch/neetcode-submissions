class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        m = [0]
        i = 0
        while i < len(prices) -1:
            m.append(max(max(prices[i+1:])-prices[i],0))
            i += 1

        return max(m)