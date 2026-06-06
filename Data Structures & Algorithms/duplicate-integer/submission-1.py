class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        h = False
        nums.sort()

        for i in range(len(nums)-1):
            if nums[i] == nums[i+1]:
                h = True
                break
        
        return h
        
