package comparator

import (
	"github.com/shushcat/rlat/text"
	"reflect"
	"sort"
	"fmt"
)

type Text = text.Text

type WordCluster map[string]int

type Comparator struct {
	density            int
	SimilarClusters    [][2]WordCluster
	Source             Text
	sourceWordClusters []WordCluster
	Target             Text
	targetWordClusters []WordCluster
}

func InitComparator(targetPath string, sourcePath string, minSharedWords int, ordered bool, window int, minWordLen int, stemming bool, stopPath string, editDist int) Comparator {
	c := Comparator{}
	c.Target = text.InitText(targetPath)
	c.Source = text.InitText(sourcePath)
	// fmt.Println(c.Target)
	var stopwords []string
	if stopPath != "" {
		stopwords = text.ParseFile(stopPath)
	}
	if editDist != 0 {
		// TODO Invoke the damlev pruning
	}
	if stemming {
		c.Target.StemWordArray()
		c.Source.StemWordArray()
	}
	// TODO Merge initialization of source and target word clusters.
	// c.targetWordClusters = initWordClusters(target, source, window,
	// minSharedWords, minWordLen, stopwords, editDist)
	// c.sourceWordClusters = initWordClusters(source, target, window,
	// minSharedWords, minWordLen, stopwords, editDist)

	c.initWordClusters(window, minSharedWords, minWordLen, stopwords, editDist)
	// c.SimilarClusters = c.initSimilarClusters(minSharedWords, ordered)
	c.initSimilarClusters(minSharedWords, ordered)
	if stemming {
		c.Target.UnstemWordArray()
		c.Source.UnstemWordArray()
	}
	// c.density = comparator_density
	// c.density()
	return c
}

// Given a comparator, returns its "density" as the ratio of sums of the
// character counts of matching words in similar clusters in the source text to
// the number of characters in the source text as a whole.

// func (c *Comparator) density() int {
// 	uniqued_hash := make(map...
// 	for _, cluster_pair := range c.SimilarClusters {
// 		for key, _ := range cluster_pair[1] {
// 			uniqued_hash[key] = [uniqued_hash[key], cluster_pair[1][key]] .flatten.uniq.compact
// 		}
// 	}
// 	numer := 0
// 	// TODO Ensure denom length is a float?
// 	denom = len(strings.Join(strings.Split(c.Source.word_array, ""))
// 	for key, _ := range uniqued_hash.keys {
// 		numer += len(strings.Split(key, "")) * len(uniqued_hash[key])
// 	}
// 	return numer / demon
// }

func indexClusters(text1 Text, text2 Text, window int, minWordLen int, stopwords []string, editDist int, intersection []string) [][]int {
	if len(stopwords) != 0 {
		// TODO Read in stopwords; split into []string.
		// intersection = reject(intersection, stopwords)
	}
	// The keys for intersectionHash are individual words, and the values
	// are the indices at which those words appear in text1.
	intersectionHash := selectKeys(text1.WordHash, intersection)
	fmt.Print(len(intersectionHash), len(text1.WordHash), len(intersection))
	// `indices` should contain all the indices, in order.
	var indices []int
	for _, i := range intersectionHash {
		indices = append(indices, i...)
	}
	// FIXME The indices listed by the following call to Println omit precisely
	// those which are not capitalized in the final report.
	sort.Ints(indices)
	fmt.Print(" ", indices[:6], "\n")
	indexClusters := partitionIndexClusters(indices, window)
	return indexClusters
}

// partitionIndexClusters groups indices that fall within window into slices.
func partitionIndexClusters(indices []int, window int) [][]int {
	i := 0
	var indexClusters [][]int
	for len(indices) > 1 {
		if (indices[i+1] - indices[i]) > window {
			indexClusters = append(indexClusters, indices[:i+1])
			indices = indices[i+1:]
			i = -1
		} else if indices[i+1] == indices[len(indices)-1] {
			indexClusters = append(indexClusters, indices[:i+2])
			break
		}
		i = i + 1
	}
	return indexClusters
}

func (c *Comparator) initSimilarClusters(minSharedWords int, ordered bool) {
	similar := c.getSharedWordClusters(minSharedWords)
	// similar := getSharedWordClusters(minSharedWords)
	if ordered == true {
		similar = getSharedOrderingClusters(similar)
	}
	c.SimilarClusters = similar
}

func (c *Comparator) initWordClusters(window int, minSharedWords int,
	minWordLen int, stopwords []string, editDist int) {
	text1 := c.Source
	text2 := c.Target
	intersection := uniqSharedWords(&text1, &text2, minWordLen, editDist, stopwords)
	ics1 := indexClusters(text1, text2, window, minWordLen, stopwords, editDist, intersection)
	fmt.Println(ics1, len(ics1))
	ics2 := indexClusters(text2, text1, window, minWordLen, stopwords, editDist, intersection)
	// fmt.Println(ics2, len(ics2))
	var wcs []WordCluster
	wcs = populateWordCluster(text1, ics1, minSharedWords)
	c.sourceWordClusters = wcs
	// fmt.Println(c.sourceWordClusters)
	wcs = populateWordCluster(text2, ics2, minSharedWords)
	c.targetWordClusters = wcs
	// fmt.Println(c.targetWordClusters)
}

func populateWordCluster(t Text, ics [][]int, minSharedWords int) (wcs []WordCluster) {
	for _, group := range ics {
		words := make(WordCluster)
		for index, _ := range group {
			words[t.WordArray[index]] = index
		}
		if (len(words) >= minSharedWords) && !(includes(wcs, words)) {
			wcs = append(wcs, words)
		}
	}
	return wcs
}

// func initWordClusters(text1 Text, text2 Text, window int, minSharedWords int, minWordLen int, stopwords []string, editDist int) []WordCluster {
// 	var wcs []WordCluster
// 	ics := indexClusters(text1, text2, window, minWordLen, stopwords, editDist)
// 	for _, group := range ics {
// 		words := make(WordCluster)
// 		for index, _ := range group {
// 			words[text1.WordArray[index]] = index
// 		}
// 		if (len(words) >= minSharedWords) && !(includes(wcs, words)) {
// 			wcs = append(wcs, words)
// 		}
// 	}
// 	return wcs
// }

// Accepts data of the form [[{}, {}], [{}, {}]...]
func getSharedOrderingClusters(sharedWordClusters [][2]WordCluster) [][2]WordCluster {
	sharedOrderingClusters := [][2]WordCluster{}
	for _, pair := range sharedWordClusters {
		intersection := intersection(pair[0], pair[1])
		// Make arrays of keys in the order they occur in the texts
		// a = pair[0].select { |x| pair[1].key?(x) }
		// .sort_by { |_, y| y }.flatten.select.with_index { |_, i| i.even? }
		// Collect key-value pairs shared by pair[0] and pair[1] into h1
		// and h2.
		h1 := make(map[string]int, len(intersection))
		for word, index := range pair[0] {
			if intersection[word] {
				h1[word] = index
			}
		}
		h2 := make(map[string]int, len(intersection))
		for word, index := range pair[1] {
			if intersection[word] {
				h2[word] = index
			}
		}
		// s1 and s2 contain the keys from h1 and h2, sorted according
		// to their values.
		s1 := sortedKeys(h1)
		s2 := sortedKeys(h2)
		if reflect.DeepEqual(s1, s2) {
			sharedOrderingClusters = append(sharedOrderingClusters, pair)
		}
	}
	return sharedOrderingClusters
}

// TODO Optimize this if possible---or, anyway, profile it after things are
// working.
func (c *Comparator) getSharedWordClusters(minSharedWords int) [][2]WordCluster {
	sharedWordClusters := [][2]WordCluster{}
	for _, sourceCluster := range c.sourceWordClusters {
		for _, targetCluster := range c.targetWordClusters {
			intersection := intersection(sourceCluster, targetCluster)
			if len(intersection) >= minSharedWords {
				// fmt.Println(intersection)
				i := make(map[string]int)
				for word, index := range targetCluster {
					if intersection[word] {
						i[word] = index
					}
				}
				j := make(map[string]int)

				for word, index := range sourceCluster {
					if intersection[word] {
						j[word] = index
					}
				}
				pair := [2]WordCluster{i, j}
				sharedWordClusters = append(sharedWordClusters, pair)
			}
		}
	}
	return sharedWordClusters
}

func filterStopWords(clust WordCluster, stopwords []string) WordCluster {
	return clust
}
