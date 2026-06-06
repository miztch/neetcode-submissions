func hasDuplicate(nums []int) bool {
    t := make(map[int]int)
    var r bool

    for _, n := range nums {
        if t[n] == 0 {
            t[n] = 1
        } else if t[n] == 1 {
            r = true
            break
        }
    }
    return r
}
