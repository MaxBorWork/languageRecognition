package main

type (
	Ngram struct {
		Name      string
		Frequency float32
		PositionDiffer int
	}

	Letter struct {
		Name string
		Frequency float32
	}

	DocsNgramCompare struct {
		TestDocTitle 	string
		TestDocLanguage string
		DocTitle     	string
		Ngrams       	[]Ngram
		Distance     	int
	}

	DocsAlphabetCompare struct {
		TestDocTitle 	string
		TestDocLanguage string
		DocTitle    	string
		Alphabet        []Letter
		Ratio    	 	float32
	}

	TestDocument struct {
		Title    string
		Language string
		Link     string
		Text     string
		Ngramms  []Ngram
		Alphabet []Letter
	}

	Document struct {
		Title     string
		Link      string
		Language  string
		ShortText string
		Ngrams    []Ngram
		Alphabet  []Letter
	}
)
