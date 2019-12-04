package main

import (
	"fmt"
	"sort"
	"strings"
)

var colOfNgrams map[string]int
var colOfLetters map[string]int

func wordToNgrams(word string) {
	wordArr := strings.Split(word, "")
	if len(wordArr) < 3 {
		return
	}

	for i := 0; i <= len(wordArr) - NgrammSize; i++ {
		var ngram string
		if strings.Contains(word, "é") || strings.Contains(word, "è") {
			ngram = word[i:i+NgrammSize+1]
			runes := []rune(ngram)
			if runes[0] == rune(65533) {
				ngram = string(runes[1:])
			} else if runes[len(runes)-1] == rune(65533) {
				ngram = string(runes[:len(runes)-1])
			}
		} else {
			ngram = word[i:i+NgrammSize]
		}

		if colOfNgrams[ngram] != 0 {
			colOfNgrams[ngram] = colOfNgrams[ngram]+1
			continue
		}

		colOfNgrams[ngram] = 1
	}
}

func processDocumentNgram(document *Document) []DocsNgramCompare {
	var resultsArr []DocsNgramCompare
	for _, testDoc := range testDocs {
		result := compareDocsNgrams(*document, testDoc)
		resultsArr = append(resultsArr, result)
	}

	return resultsArr
}

func compareDocsNgrams(doc Document, testDoc TestDocument) DocsNgramCompare {
	var ngramsWithDifference []Ngram
	var distance int
	testDocNgrams := testDoc.Ngramms[:len(doc.Ngrams)-1]
	for i, ngram := range doc.Ngrams {
		ngramWithDiff := Ngram{
			Name:           ngram.Name,
			Frequency:      ngram.Frequency,
		}

		if ngram.Name == "ièc" {
			fmt.Println("FOUND")
		}

		for j, testNgram := range testDocNgrams {
			if testNgram.Name == ngram.Name {
				if ngram.Name == "ièc" {
					fmt.Println("FOUND")
				}
				ngramWithDiff.PositionDiffer = Abs(i-j)
				break
			}
		}

		if ngramWithDiff.PositionDiffer == 0 {
			ngramWithDiff.PositionDiffer = len(doc.Ngrams)
		}

		ngramsWithDifference = append(ngramsWithDifference, ngramWithDiff)
		distance += ngramWithDiff.PositionDiffer
	}

	result := DocsNgramCompare{
		TestDocTitle:	 	testDoc.Title,
		TestDocLanguage:	testDoc.Language,
		DocTitle:       	doc.Title,
		Ngrams: 			ngramsWithDifference,
		Distance: 			distance,
	}
	return result
}

func getComparisonWithMinimalDistance(resultsArr []DocsNgramCompare) DocsNgramCompare {
	sort.Slice(resultsArr, func(i, j int) bool {
		return resultsArr[i].Distance < resultsArr[j].Distance
	})

	return resultsArr[0]
}

func getTestDoc(title string) TestDocument {
	for _, testDoc := range testDocs {
		if testDoc.Title == title {
			return testDoc
		}
	}

	return TestDocument{}
}

func textToNgrams(text string) []Ngram {
	wordsArr := splitText(text)

	colOfNgrams = make(map[string]int)
	for _, word := range wordsArr {
		word = clearWord(word)
		wordToNgrams(word)
	}

	return createNgramsArrForDocument(len(wordsArr))
}

func createNgramsArrForDocument(docColOfWords int) []Ngram {
	var ngrams []Ngram
	for ngramName, col := range colOfNgrams {
		ngram := Ngram{
			Name:      ngramName,
			Frequency: float32(col) / float32(docColOfWords),
		}
		ngrams = append(ngrams, ngram)
	}

	sort.Slice(ngrams, func(i, j int) bool {
		return ngrams[i].Frequency > ngrams[j].Frequency
	})

	return ngrams
}

func getDocLanguage(title string) string {
	testDoc := getTestDoc(title)

	if testDoc.Language != "" {
		return testDoc.Language
	}

	return ""
}
