class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        lowest = prices[0]
        income = 0

        for i,p in enumerate(prices):
            if lowest > p:
                lowest = p
                continue
            else:
                if income < p - lowest:
                    income = p - lowest

        return income