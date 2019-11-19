package main

import (
	"sort"
	"strings"
)

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

func processDocument(document Document) string {
	var resultsArr []DocsCompare
	for _, testDoc := range testDocs {
		result := compareDocs(document, testDoc)
		resultsArr = append(resultsArr, result)
	}

	testDocTitle := getTestDocWithMinimalDistance(resultsArr)
	testDoc := getTestDoc(testDocTitle)

	if testDoc.Language != "" {
		return testDoc.Language
	}

	return ""
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
				ngramWithDiff.PositionDiffer = i-j
				break
			}
		}

		ngramWithDiff.PositionDiffer = len(doc.Ngrams)
		ngramsWithDifference = append(ngramsWithDifference, ngramWithDiff)
		distance += ngramWithDiff.PositionDiffer
	}

	result := DocsCompare{
		TestDocTitle:	testDoc.Title,
		DocTitle:       doc.Title,
		Ngrams: 		ngramsWithDifference,
		Distance: 		distance,
	}
	return result
}

func getTestDocWithMinimalDistance(resultsArr []DocsCompare) string {
	sort.Slice(resultsArr, func(i, j int) bool {
		return resultsArr[i].Distance < resultsArr[j].Distance
	})

	return resultsArr[0].TestDocTitle
}

func getTestDoc(title string) TestDocument {
	for _, testDoc := range testDocs {
		if testDoc.Title == title {
			return testDoc
		}
	}

	return TestDocument{}
}
