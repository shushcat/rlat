package comparator

import (
	"reflect"
	"github.com/shushcat/rlat/damlev"
	"sort"
)

// type wordClusters []map[string]int

type set interface {
	intersection() set
}

// The `includes` function checks whether a slice `sl` contains an element `e`.

// type set interface {
// 	includes() bool
// }

// func includes(sl []string, e string) bool {
// 	for _, v := range sl {
// 		if v == e {
// 			return true
// 		}
// 	}
// 	return false
// }

// func includes(wcs []WordCluster, words map[string]int) bool {
func includes(wcs []WordCluster, wc1 WordCluster) bool {
	for _, wc2 := range wcs {
		if (reflect.DeepEqual(wc2, wc1)) {
			return true
		}
	}
	return false
}

func wcIntersection(wc1 WordCluster, wc2 WordCluster) map[string]bool {
	intersection := make(map[string]bool)
	for word, _ := range wc1 {
		if _, ok := wc2[word]; ok {
			intersection[word] = true
		}
	}
	return intersection
}

func selectKeys(hash1 map[string][]int, keys []string) map[string][]int {
	hash2 := make(map[string][]int, len(keys))
	for k, v := range hash1 {
		for i:=0; i<len(keys); i++ {
			if k == keys[i] {
				hash2[k] = v
			}
		}
	}
	return hash2
}

// The `uniqSharedWords` function takes pointers to two texts, a minimum word
// length, and a permissible Damarau-Levenshtein edit distance.  Unique words
// that are shared by the two texts, meet the length and edit distance
// requirements, and are not in the stopwords list are returned as a string
// slice.
func uniqSharedWords(t1 *Text, t2 *Text, minWordLen int, editDist int, stopwords []string) []string {
	var shared []string
	uniqueShared := make(map[string]bool)
	if editDist > 0 {
		for _, word1 := range t1.WordArray {
			for _, word2 := range t2.WordArray {
				if editDist >= damlev.Distance(word1, word2) {
					uniqueShared[word1] = true
					uniqueShared[word2] = true
				}
			}
		}
	} else {
		for _, word1 := range t1.WordArray {
			for _, word2 := range t2.WordArray {
				if word1 == word2 {
					uniqueShared[word1] = true
				}
			}
		}
	}
	// Reject stopwords.
	for _, v := range stopwords {
		if _, ok := uniqueShared[v]; ok {
			delete(uniqueShared, v)
		}
	}
	// Reject words shorter than minWordLen from the intersection.
	for k := range uniqueShared {
		if len(k) < minWordLen {
			delete(uniqueShared, k)
		} else {
			shared = append(shared, k)
		}
	}
	return shared
}

// reject returns the elements of the string slice s1 that are not present in
// the string slice s2.
func reject(s1 []string, s2 []string) []string {
	reject := make(map[string]bool)
	for _, v := range s2 {
		reject[v] = true
	}
	n := 0
	for _, w := range s1 {
		if !reject[w] {
			s1[n] = w
			n++
		}
	}
	return s1[:n]
}

// The `sortedKeys` function returns a passed WordCluster's words as a string
// slice, ordered by the words indices.
func sortedKeys(wc WordCluster) (s []string) {
	type kv struct {
		key   int
		value string
	}
	// Extract the key-value pairs into a slice of kv structs.
	var kvs []kv
	for word, indices := range wc {
		for _, i := range indices {
			kvs = append(kvs, kv{i, word})
		}
	}
	// Sort the slice of kv structs based on keys, from least to
	// greatest.
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].key < kvs[j].key
	})
	for _, kv := range kvs {
		s = append(s, kv.value)
	}
	return s
}

func (wc WordCluster) FlatValues() []int {
	valMap := make(map[int]bool)
	var valFlat []int
	for _, ic := range wc {
		for _, i := range ic {
			valMap[i] = true
		}
	}
	for i, _ := range valMap {
		valFlat = append(valFlat, i)
	}
	return valFlat
}
