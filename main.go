package main

import (
	"github.com/gin-gonic/gin"
	"sort"
)

const (
	PORT = "3333"
	NgrammSize = 3
)

var langColOfWords int
var testDocs []TestDocument
var docsMap map[string]Document

func init() {
	colOfNgramm = make(map[string]int)

	prepareNgrams("English")

	docsMap = make(map[string]Document)

	docsMap["chess_fr"] =  Document{
		Title:     "Échecs",
		Language:  "fr",
		Link:      "https://fr.wikipedia.org/wiki/%C3%89checs",
		ShortText: "",
	}

	docsMap["saint_martin_fr"] = Document{
		Title:     "Saint-Martin-sur-Écaillon",
		Language:  "fr",
		Link:      "https://fr.wikipedia.org/wiki/Saint-Martin-sur-%C3%89caillon",
		ShortText: "",
	}

	docsMap["chretien_fr"] = Document{
		Title:     "Albert_Chrétien",
		Language:  "fr",
		Link:      "https://fr.wikipedia.org/wiki/Albert_Chr%C3%A9tien",
		ShortText: "",
	}

	docsMap["santos_fr"] = Document{
		Title:     "Alberto_Santos-Dumont",
		Language:  "fr",
		Link:      "https://fr.wikipedia.org/wiki/Alberto_Santos-Dumont",
		ShortText: "",
	}

	docsMap["soprano_fr"] = Document{
		Title:     "Mezzo-soprano",
		Language:  "fr",
		Link:      "https://fr.wikipedia.org/wiki/Mezzo-soprano",
		ShortText: "",
	}


	docsMap["futurama_en"] = Document{
		Title:     "Futurama",
		Language:  "en",
		Link:      "https://en.wikipedia.org/wiki/Futurama",
		ShortText: "",
	}

	docsMap["weavers_needle_en"] = Document{
		Title:     "Weavers_Needle",
		Language:  "en",
		Link:      "https://en.wikipedia.org/wiki/Weavers_Needle",
		ShortText: "",
	}

	docsMap["twisted_sister_en"] = Document{
		Title:     "Twisted_Sister",
		Language:  "en",
		Link:      "https://en.wikipedia.org/wiki/Twisted_Sister",
		ShortText: "",
	}

	docsMap["new_york_en"] = Document{
		Title:     "New_York_City",
		Language:  "en",
		Link:      "https://en.wikipedia.org/wiki/New_York_City",
		ShortText: "",
	}

	docsMap["Buffalo_Oklahoma_en"] = Document{
		Title:     "Buffalo,_Oklahoma",
		Language:  "en",
		Link:      "https://en.wikipedia.org/wiki/Buffalo,_Oklahoma",
		ShortText: "",
	}
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
	route.GET("/{title}/ngram", NgramMethod)
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

func createNgramsArrForDocument() []Ngram {
	var ngrams []Ngram
	for ngramName, col := range colOfNgramm {
		ngram := Ngram{
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
