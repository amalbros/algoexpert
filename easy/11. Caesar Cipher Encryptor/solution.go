package program

import "strings"

func CaesarCipherEncryptor(str string, key int) string {
	// Write your code here.
	runes := []rune(str) // converting into ASCII numbers
	alphabets := "abcdefghijklmnopqrstuvwxyz"
	for i, char := range runes {
		index := strings.Index(alphabets, string(char))
		if index == -1 {
			return ""
		}
		new_index := (index + key) % 26
		runes[i] = rune(alphabets[new_index])
	}
	return string(runes)
}
