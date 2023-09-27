package Leetcode

func MergeAlternately(word1 string, word2 string) (word3 string) {
	var i = 0
	for i != len(word1) && i != len(word2) {
		word3 = word3 + string(word1[i]) + string(word2[i])
		i++
	}
	if len(word1) > len(word2) {
		for i = len(word2); i < len(word1); i++ {
			word3 = word3 + string(word1[i])
		}
	} else {
		for i = len(word1); i < len(word2); i++ {
			word3 = word3 + string(word2[i])
		}
	}
	return
}
