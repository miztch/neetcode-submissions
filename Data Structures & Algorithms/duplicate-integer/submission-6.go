import "slices"
func hasDuplicate(nums []int) bool {
    // t := make(map[int]bool)

    // for _, n := range nums {
    //     if t[n] {
    //         return true
    //     }
    //     t[n] = true
    // }
    // return false
    sz := len(nums)

    if sz < 2 {
        return false
    }
    
    slices.Sort(nums)
    for i := 0; i < sz-1; i++ {
        if nums[i] == nums[i+1] {
            return true
        }
    }
    return false
}
