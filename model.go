package main

type (
	Ngram struct {
		Name      string
		Frequency float32
		PositionDiffer int
	}

	DocsCompare struct {
		TestDocTitle string
		TestDocLanguage string
		DocTitle     string
		Ngrams       []Ngram
		Distance     int
	}

	TestDocument struct {
		Title    string
		Language string
		Link     string
		Text     string
		Ngramms  []Ngram
	}

	Document struct {
		Title     string
		Link      string
		Language string
		ShortText string
		Ngrams    []Ngram
	}
)
