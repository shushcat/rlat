package stemmer

import (
	"testing"
)

func TestStemmerPresent (t *testing.T) {
	word := "addressable"
	if !(stem(word) == "address") {
		t.Errorf("stemming '%v' did not yield 'address'", word)
	}
}
