package comparator

import (
	"github.com/shushcat/rlat/text"
	"reflect"
	"sort"
)

type Text = text.Text

type WordCluster map[string][]int

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
	var stopwords []string
	if stopPath != "" {
		stopwords = text.ParseFile(stopPath)
	}
	if editDist != 0 {
		// TODO Invoke the DamLev pruning?
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
	// fmt.Print(len(intersectionHash), len(text1.WordHash), len(intersection))
	// `indices` should contain all the indices, in order.
	var indices []int
	for _, i := range intersectionHash {
		indices = append(indices, i...)
	}
	sort.Ints(indices)
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
	ics2 := indexClusters(text2, text1, window, minWordLen, stopwords, editDist, intersection)
	// FIXME WordClusters are injective, but they should
	// associate words with sets of indices.
	var wcs []WordCluster
	wcs = populateWordCluster(text1, ics1, minSharedWords)
	// return
	c.sourceWordClusters = wcs
	// fmt.Println(c.sourceWordClusters)
	wcs = populateWordCluster(text2, ics2, minSharedWords)
	c.targetWordClusters = wcs
	// fmt.Println(c.targetWordClusters)
}

func populateWordCluster(t Text, ics [][]int, minSharedWords int) (wcs []WordCluster) {
	for _, ic := range ics {
		wc := make(WordCluster)
		for i, _ := range ic {
			wc[t.WordArray[i]] = append(wc[t.WordArray[i]], i)
		}
		if (len(wc) >= minSharedWords) && !(includes(wcs, wc)) {
			wcs = append(wcs, wc)
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

// The `getSharedOrderingClusters` function returns those WordCluster pairs in
// which the shared words occur in the same order.
func getSharedOrderingClusters(sharedWordClusters [][2]WordCluster) [][2]WordCluster {
	sharedOrderingClusters := [][2]WordCluster{}
	for _, wcPair := range sharedWordClusters {
		intersection := wcIntersection(wcPair[0], wcPair[1])
		// Make arrays of keys in the order they occur in the texts
		// a = wcPair[0].select { |x| wcPair[1].key?(x) }
		// .sort_by { |_, y| y }.flatten.select.with_index { |_, i| i.even? }
		// Collect key-value pairs shared by wcPair[0] and wcPair[1] into h1
		// and h2.
		// h1 := make(map[string]int, len(intersection))
		h1 := make(map[string][]int)
		for word, is := range wcPair[0] {
			if _, ok := intersection[word]; ok {
				h1[word] = is
			}
		}
		h2 := make(map[string][]int, len(intersection))
		for word, is := range wcPair[1] {
			if _, ok := intersection[word]; ok {
				h2[word] = is
			}
		}
		// s1 and s2 contain the keys from h1 and h2, sorted according
		// to their values.
		s1 := sortedKeys(h1)
		s2 := sortedKeys(h2)
		if reflect.DeepEqual(s1, s2) {
			sharedOrderingClusters = append(sharedOrderingClusters, wcPair)
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
			intersection := wcIntersection(sourceCluster, targetCluster)
			if len(intersection) >= minSharedWords {
				i := make(WordCluster)
				for word, indices := range targetCluster {
					if intersection[word] {
						i[word] = indices
					}
				}
				j := make(WordCluster)
				for word, indices := range sourceCluster {
					if intersection[word] {
						j[word] = indices
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
