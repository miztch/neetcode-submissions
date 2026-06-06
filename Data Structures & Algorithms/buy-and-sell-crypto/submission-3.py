class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        lowest = prices[0]
        income = 0

        for i,p in enumerate(prices):
            lowest = min(lowest, p)
            income = max(income, p - lowest)

        return income