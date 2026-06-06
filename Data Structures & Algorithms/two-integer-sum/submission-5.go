func twoSum(nums []int, target int) []int {
   seen := make(map[int]int)
   for i,n := range(nums) {
        c := target - n
        if value, ok := seen[c]; ok {
            return []int{value,i}
        }

        seen[n] = i
    } 
    return nil
}
