package main

type (
	Ngramm struct {
		Name      string
		Frequency float32
	}

	TestDocument struct {
		Title    string
		Language string
		Link     string
		Text     string
		Ngramms  []Ngramm
	}

	Document struct {
		Title     string
		Link      string
		ShortText string
		Ngrams    []Ngramm
	}
)
