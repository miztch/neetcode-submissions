func maxProfit(prices []int) int {
    l := prices[0]
    i := 0

    for _,p := range(prices) {
        l = min(l, p)
        i = max(i, p-l)
    }

    return i
}
