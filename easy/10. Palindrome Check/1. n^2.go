package program

func IsPalindrome(str string) bool {
	// Write your code here.
	rev_str := make([]byte, len(str))
	for i := len(str) - 1; i >= 0; i-- {
		j := len(str) - i - 1
		rev_str[j] = str[i]
	}
	for i := range str {
		if rev_str[i] != str[i] {
			return false
		}
	}
	return true
}
