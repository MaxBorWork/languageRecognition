package main

import (
	"sort"
	"strings"
)

func createLettersArrForDocument(lettersInText int) []Letter {
	var letters []Letter
	for ngramName, col := range colOfLetters {
		letter := Letter{
			Name:      ngramName,
			Frequency: float32(col) / float32(lettersInText),
		}
		letters = append(letters, letter)
	}

	return letters
}

func textToLetters(text string) []Letter {
	var lettersInText int
	wordsArr := splitText(text)

	colOfLetters = make(map[string]int)
	for _, word := range wordsArr {
		word = clearWord(word)
		lettersInText += wordToLetters(word)
	}

	return createLettersArrForDocument(lettersInText)
}

func wordToLetters(word string) int {
	var lettersInWord int
	wordArr := strings.Split(word, "")
	for _, letter := range wordArr {
		if letter == "'" || letter == "-" || letter == "[" || letter == "]" {
			continue
		}
		lettersInWord = lettersInWord+1
		if colOfLetters[letter] != 0 {
			colOfLetters[letter] = colOfLetters[letter]+1
			continue
		}

		colOfLetters[letter] = 1
	}

	return lettersInWord
}

func processDocumentAlphabet(document *Document) []DocsAlphabetCompare {
	var resultsArr []DocsAlphabetCompare
	for _, testDoc := range testDocs {
		result := compareDocsAlphabets(*document, testDoc)
		resultsArr = append(resultsArr, result)
	}

	return resultsArr
}

func compareDocsAlphabets(doc Document, testDoc TestDocument) DocsAlphabetCompare {
	var alpabetsIntersection []Letter
	for _, letter := range doc.Alphabet {
		for _, testLetter := range testDoc.Alphabet {
			if testLetter.Name == letter.Name {
				alpabetsIntersectionElement := Letter{
					Name:      letter.Name,
					Frequency: 0,
				}

				alpabetsIntersection = append(alpabetsIntersection, alpabetsIntersectionElement)
				break
			}
		}
	}

	ratio := float32(len(alpabetsIntersection)) / float32(len(doc.Alphabet))

	result := DocsAlphabetCompare{
		TestDocTitle:	 	testDoc.Title,
		TestDocLanguage:	testDoc.Language,
		DocTitle:       	doc.Title,
		Alphabet: 			alpabetsIntersection,
		Ratio: 				ratio,
	}
	return result
}

func getComparisonWithMaxRatio(resultsArr []DocsAlphabetCompare) DocsAlphabetCompare {
	sort.Slice(resultsArr, func(i, j int) bool {
		return resultsArr[i].Ratio > resultsArr[j].Ratio
	})

	return resultsArr[0]
}
