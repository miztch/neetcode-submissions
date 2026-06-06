class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i in range(len(nums)):
            c = target - nums[i]
            if not c in nums:
                continue

            if nums[i] == c:
                if nums.count(c) == 2:
                    return [i, nums.index(c,i+1)]
                else:
                    continue
            else:
                return [i, nums.index(c)]
                break
