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

func includes(wcs []WordCluster, words map[string]int) bool {
	for _, cluster := range wcs {
		if (reflect.DeepEqual(cluster, words)) {
			return true
		}
	}
	return false
}

func intersection(h1 map[string]int, h2 map[string]int) map[string]bool {
	intersection := make(map[string]bool)
	for word, _ := range h1 {
		if _, ok := h2[word]; ok {
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

// The `sharedWords` function takes pointers to two texts, a minimum word
// length, and a permissible Damarau-Levenschtein edit distance.  Words 
// that are shared by the two texts, meet the length and edit distance
// requirements, and are not in the stopwords list are returned as a string
// slice.
func sharedWords(t1 *Text, t2 *Text, minWordLen int, editDist int, stopwords []string) []string {
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

// sortedKeys accepts a map of strings to integers and returns the key strings
// as a slice sorted from least to greatest by their integer values.
func sortedKeys(m map[string]int) (s []string) {
	// kv holds key-value pairs.
	type kv struct {
		key   string
		value int
	}
	// Extract the key-value pairs into a slice of kv structs.
	var kvs []kv
	for k, v := range m {
		kvs = append(kvs, kv{k, v})
	}
	// Sort the slice of kv structs based on values, from least to
	// greatest.
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].value < kvs[j].value
	})
	for _, kv := range kvs {
		s = append(s, kv.key)
	}
	return s
}

func (clust WordCluster) Values() []int {
	var vs []int
	for _, v := range clust {
		vs = append(vs, v)
	}
	return vs
}
