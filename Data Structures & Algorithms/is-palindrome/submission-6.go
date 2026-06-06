// import (
//     "strings"
//     "unicode"
// )

func isAlnum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func isPalindrome(s string) bool {
	l := 0
    r := len(s) - 1

    // rs := []rune(s) // makes space complexity O(n)
    for l < r {
        for l < r && !isAlnum(rune(s[l])) {
            l += 1
        }

        for l < r && !isAlnum(rune(s[r])) {
            r -= 1
        }

        if strings.ToLower(string(rune(s[l]))) != strings.ToLower(string(rune(s[r]))) {
            return false
        }

        l += 1
        r -= 1
    }

    return true
}
