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

    rs := []rune(s)
    for (l < r) {
        for (l < r && !isAlnum(rs[l])) {
            l += 1
        }

        for (l < r && !isAlnum(rs[r])) {
            r -= 1
        }

        if strings.ToLower(string(rs[l])) != strings.ToLower(string(rs[r])) {
            return false
            break
        }

        l += 1
        r -= 1
    }

    return true
}
