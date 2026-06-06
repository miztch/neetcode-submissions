import "maps"

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    m := make(map[byte]int)
    n := make(map[byte]int)
    for i := 0; i < len(s);i++ {
        m[s[i]] += 1;
        n[t[i]] += 1;
    }
    return maps.Equal(m,n)
}
