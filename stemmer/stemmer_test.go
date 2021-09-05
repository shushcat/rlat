package stemmer

import "testing"

func TestStemmer(t *testing.T) {
	var wordstem, word string
	wordstem, word = Stem("addressable")
	if wordstem != "address" {
		t.Errorf("stemming '%v' yielded '%v', not 'address'", word, wordstem)
	}
	wordstem, word = Stem("loving")
	if wordstem != "love" {
		t.Errorf("stemming '%v' yielded '%v', not 'love'", word, wordstem)
	}
	wordstem, word = Stem("Argumentative")
	if wordstem != "argument" {
		t.Errorf("stemming '%v' yielded '%v', not 'argument'", word, wordstem)
	}
	wordstem, word = Stem("astonishment")
	if wordstem != "astonish" {
		t.Errorf("stemming '%v' yielded '%v', not 'astonish'", word, wordstem)
	}
}
