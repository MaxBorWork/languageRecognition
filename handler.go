package main

import (
	"github.com/gin-gonic/gin"
	"mediawiki"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func NgramMethod(c *gin.Context) {
	var text string
	id := c.Param("title")
	document := docsMap[id]
	url := mediawiki.CreateUrl(document.Language, document.Title)
	text = mediawiki.MediaWikiRequest(url)
	document.ShortText = text
	document.Ngrams = textToNgrams(text)
	resultsArr := processDocumentNgram(&document)

	comparison := getComparisonWithMinimalDistance(resultsArr)
	language := getDocLanguage(comparison.TestDocTitle)
	docsMap[id] = document
	if language != "" {
		c.HTML(http.StatusOK, "result_ngrams.html", gin.H {
			"Id": id,
			"Title" : document.Title,
			"Link": document.Link,
			"Language": language,
			"Ngrams": comparison.Ngrams,
			"ResultsArr": resultsArr,
		})

		return
	}

	c.String(http.StatusNotFound, "Can't detect doc language")
}

func AlphabetMethod(c *gin.Context)  {
	var text string
	id := c.Param("title")
	document := docsMap[id]
	url := mediawiki.CreateUrl(document.Language, document.Title)
	text = mediawiki.MediaWikiRequest(url)
	document.ShortText = text
	docsMap[id] = document
	document.Alphabet = textToLetters(text)
	resultsArr := processDocumentAlphabet(&document)

	comparison := getComparisonWithMaxRatio(resultsArr)
	language := getDocLanguage(comparison.TestDocTitle)
	docsMap[id] = document
	if language != "" {
		c.HTML(http.StatusOK, "result_alphabet.html", gin.H{
			"Id": id,
			"Title" : document.Title,
			"Link": document.Link,
			"Language": language,
			"Alphabet": comparison.Alphabet,
			"Ratio": comparison.Ratio,
			"ResultsArr": resultsArr,
		})

		return
	}

	c.String(http.StatusNotFound, "Can't detect doc language")
}

func File(c *gin.Context) {
	id := c.Param("title")
	document := docsMap[id]
	c.HTML(http.StatusOK, "file.html", gin.H{
		"Title" : document.Title,
		"Link": document.Link,
		"Text": document.ShortText,
	})
}
