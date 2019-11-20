package main

import (
	"sort"
	"strings"
)

var colOfNgrams map[string]int

func wordToNgrams(word string) {
	wordArr := strings.Split(word, "")
	if len(wordArr) < 3 {
		return
	}

	for i := 0; i <= len(wordArr) - NgrammSize; i++ {
		ngram := word[i:i+NgrammSize]
		if colOfNgrams[ngram] != 0 {
			colOfNgrams[ngram] = colOfNgrams[ngram]+1
			continue
		}

		colOfNgrams[ngram] = 1
	}
}

func processDocument(document *Document) []DocsCompare {
	var resultsArr []DocsCompare
	for _, testDoc := range testDocs {
		result := compareDocs(*document, testDoc)
		resultsArr = append(resultsArr, result)
	}

	return resultsArr
}

func compareDocs(doc Document, testDoc TestDocument) DocsCompare {
	var ngramsWithDifference []Ngram
	var distance int
	for i, ngram := range doc.Ngrams {
		ngramWithDiff := Ngram{
			Name:           ngram.Name,
			Frequency:      ngram.Frequency,
		}

		for j, testNgram := range testDoc.Ngramms {
			if testNgram.Name == ngram.Name {
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

	result := DocsCompare{
		TestDocTitle:	 	testDoc.Title,
		TestDocLanguage:	testDoc.Language,
		DocTitle:       	doc.Title,
		Ngrams: 			ngramsWithDifference,
		Distance: 			distance,
	}
	return result
}

func getComparisonWithMinimalDistance(resultsArr []DocsCompare) DocsCompare {
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