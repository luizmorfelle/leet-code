package main

func mergeAlternately(word1 string, word2 string) string {

	len1 := len(word1)
	len2 := len(word2)

	lenght := len1

	if len2 > lenght {
		lenght = len2
	}

	res := ""

	for i := 0; i < lenght; i++ {
		if i < len1 && word1[i] != 0 {
			res += string(word1[i])
		}

		if i < len2 && word2[i] != 0 {
			res += string(word2[i])
		}
	}

	return res
}

func main() {
	println(mergeAlternately("abc", "pqr"))
}
