package lorem

import (
	golorem "github.com/drhodes/golorem"
)

// Service is a generic interface
type Service interface {
	// generate a word with at least min letters and at most max letters
	Word(min, max int) string

	// generate a sentence with at least min words and at most max words
	Sentence(min, max int) string

	// generate a paragraph with at least min words and at most max words
	Paragraph(min, max int) string
}

func NewLoremService() Service {
	return &loremService{}
}

type loremService struct{}

func (loremService) Word(min, max int) string {
	return golorem.Word(min, max)
}

func (loremService) Sentence(min, max int) string {
	return golorem.Sentence(min, max)
}

func (loremService) Paragraph(min, max int) string {
	return golorem.Paragraph(min, max)
}
