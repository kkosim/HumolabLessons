package Leetcode

func ReverseVowels(s string) string {
	var vowels []int32
	for i, val := range s {
		if HasVowel(s[i]) {
			vowels = append(vowels, val)
		}
	}
	str := []byte(s)
	k := len(vowels) - 1
	j := 0
	for i := 0; i < len(str); i++ {
		if HasVowel(s[i]) {
			str[i] = byte(vowels[k-j])
			j++
		}
	}
	return string(str)
}
func HasVowel(symbol byte) bool {
	arr := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for _, v := range arr {
		if v == symbol {
			return true
		}
	}
	return false
}
