package text

import (
	// "github.com/shushcat/rlat/damlev"
	"bufio"
	"github.com/shushcat/rlat/stemmer"
	"log"
	"os"
	"regexp"
	"strings"
)

type Text struct {
	FileName          string
	originalWordArray []string
	WordArray         []string
	WordHash          map[string][]int // indices of words
}

func InitText(path string) Text {
	t := Text{}
	t.WordHash = make(map[string][]int)
	t.FileName = path
	t.WordArray = ParseFile(path)
	t.initWordHash(t.WordArray)
	return t
}

func (t *Text) StemWordArray() {
	copy(t.originalWordArray, t.WordArray)
	for i, word := range t.WordArray {
		t.WordArray[i] = stemmer.Stem(word) //TODO Setup stemmer.
	}
}

func (t *Text) UnstemWordArray() {
	t.WordArray = t.originalWordArray
}

func (t *Text) initWordHash(wordArray []string) {
	// wordHash := make(map[string][]int)
	for i, word := range wordArray {
		t.WordHash[word] = append(t.WordHash[word], i)
	}
}

// ParseFile splits a text file into a slice of words.
// Stemming work may grow out from here.
func ParseFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	rx, err := regexp.Compile("[^a-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	var words []string
	for scanner.Scan() {
		word := rx.ReplaceAllString(strings.ToLower(scanner.Text()), "")
		words = append(words, word)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return words
}
