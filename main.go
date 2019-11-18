package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sort"
)

const (
	PORT = "3333"
	NgrammSize = 3
)

var langColOfWords int
var testDocs []TestDocument

func init() {
	colOfNgramm = make(map[string]int)

	prepareNgrams("English")

	fmt.Println(testDocs[0])
}

func main() {
	//router := route()
	//router.LoadHTMLGlob("templates/*")
	//err := router.Run(":" + PORT)
	//if err != nil {
	//	panic(err)
	//}
}

func route() *gin.Engine {
	route := gin.Default()
	route.GET("/", Index)
	route.GET("/{title}/ngramm", NgrammMethod)
	route.GET("/{title}/alphabet", AlphabetMethod)
	return route
}

func prepareNgrams(languange string) {
	dirName := "testDocs" + languange
	files := getFilesFromDirectory(dirName)
	for _, file := range files {
		parseTestDoc(languange, file.Name())
	}
}

func parseTestDoc(languange, title string) {
	path := "testDocs" + languange + "/" + title
	text, err := readFile(path)
	if err != nil {
		panic(err)
	}

	wordsArr := splitText(text)
	langColOfWords = langColOfWords + len(wordsArr)

	for _, word := range wordsArr {
		word = clearWord(word)
		wordToNgrams(word)
	}

	ngrams := createNgramsArrForDocument()

	testDocs = append(testDocs, TestDocument{
		Title:    title,
		Language: languange,
		Ngramms:  ngrams,
	})
}

func createNgramsArrForDocument() []Ngramm {
	var ngrams []Ngramm
	for ngramName, col := range colOfNgramm {
		ngram := Ngramm{
			Name:      ngramName,
			Frequency: float32(col) / float32(langColOfWords),
		}
		ngrams = append(ngrams, ngram)
	}

	sort.Slice(ngrams, func(i, j int) bool {
		return ngrams[i].Frequency > ngrams[j].Frequency
	})

	return ngrams
}
