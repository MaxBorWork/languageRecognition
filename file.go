package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func getFilesFromDirectory(name string) []os.FileInfo {
	files, err := ioutil.ReadDir(name)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func readFile(path string) (string, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(dat), nil
}

func splitText(text string) []string {
	return strings.Fields(text)
}

func clearWord(word string) string {
	word = strings.ToLower(word)

	word = strings.Replace(word, ".", "", -1)
	word = strings.Replace(word, ",", "", -1)
	word = strings.Replace(word, "!", "", -1)
	word = strings.Replace(word, "?", "", -1)
	word = strings.Replace(word, "(", "", -1)
	word = strings.Replace(word, ")", "", -1)
	word = strings.Replace(word, ";", "", -1)
	word = strings.Replace(word, ":", "", -1)
	word = strings.Replace(word, "«", "", -1)
	word = strings.Replace(word, "»", "", -1)
	word = strings.Replace(word, "...", "", -1)
	word = strings.Replace(word, "----", "", -1)
	word = strings.Replace(word, "+", "", -1)
	word = strings.Replace(word, "=", " ", -1)
	word = strings.Replace(word, "≠", "", -1)
	word = strings.Replace(word, "#", "", -1)
	word = strings.Replace(word, "\"", "", -1)
	word = strings.Replace(word, "--", "", -1)
	word = strings.Replace(word, "—", "", -1)
	word = strings.Replace(word, "‘", "", -1)
	word = strings.Replace(word, "’", "", -1)

	re := regexp.MustCompile(`[0-9]`)
	word = re.ReplaceAllString(word, "")


	return word
}
