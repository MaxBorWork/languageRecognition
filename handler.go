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
	resultsArr := processDocument(&document)

	comparison := getComparisonWithMinimalDistance(resultsArr)
	language := getDocLanguage(comparison.TestDocTitle)
	if language != "" {
		c.HTML(http.StatusOK, "result.html", gin.H{
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
	//words = workWithText(text, languageMap[document.Language])
	//c.HTML(http.StatusOK, "result.html", gin.H{
	//	"Title" : document.Title,
	//	"Link": document.Link,
	//	"Words": words,
	//})
}

//func File(c *gin.Context) {
//	c.HTML(http.StatusOK, "file.html", gin.H{
//		"Title" : document.Title,
//		"Link": document.Link,
//		"Text": document.SHortText,
//	})
//}
