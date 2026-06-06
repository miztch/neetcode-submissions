func hasDuplicate(nums []int) bool {
    t := make(map[int]bool)

    for _, n := range nums {
        if t[n] {
            return true
        }
        t[n] = true
    }
    return false
}
