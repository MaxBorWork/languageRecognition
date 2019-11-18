package main

import "strings"

var colOfNgramm map[string]int

func wordToNgrams(word string) {
	wordArr := strings.Split(word, "")
	if len(wordArr) < 3 {
		return
	}

	for i := 0; i <= len(wordArr) - NgrammSize; i++ {
		ngram := word[i:i+NgrammSize]
		if colOfNgramm[ngram] != 0 {
			colOfNgramm[ngram] = colOfNgramm[ngram]+1
			continue
		}

		colOfNgramm[ngram] = 1
	}
}
