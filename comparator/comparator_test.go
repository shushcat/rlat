package comparator

import (
	"reflect"
	"testing"
	// "github.com/shushcat/rlat/text"
)

func TestDissimilarComparator(t *testing.T) {
	son2 := "../sonnets/Sonnet II.txt"
	son5 := "../sonnets/Sonnet V.txt"
	comp := InitComparator(son2, son5, 3, true, 10, 4, false, "", 0)
	if len(comp.Source.WordArray) != 104 {
		t.Errorf("len(comp.Source.WordArray) is %d, but should be 104", len(comp.Source.WordArray))
	}
	if len(comp.Target.WordArray) != 115 {
		t.Errorf("len(comp.Target.WordArray) is %d, but should be 115", len(comp.Target.WordArray))
	}
	if len(comp.SimilarClusters) != 0 {
		t.Errorf("len(comp.SimilarClusters) is %d, but should be 0", len(comp.SimilarClusters))
	}
}

func TestIdentityComparator(t *testing.T) {
	son2 := "../sonnets/Sonnet II.txt"
	comp := InitComparator(son2, son2, 3, true, 10, 1, false, "", 0)
	// There should be only one cluster
	if len(comp.SimilarClusters) != 1 {
		t.Errorf("len(comp.SimilarClusters) is %d, but should be 1", len(comp.SimilarClusters))
	}
	// Make sure the last word of the sonnet is found
	if comp.SimilarClusters[0][0]["cold"] == nil {
		t.Errorf("The last word of Sonnet II is \"cold\", but comp.SimilarClusters[0][0][\"cold\"] is %d", comp.SimilarClusters[0][0]["cold"])
	}
	// Mark sure the cluster's elements are equal
	if !reflect.DeepEqual(comp.SimilarClusters[0][0], comp.SimilarClusters[0][1]) {
		t.Error("The pair at comp.SimilarClusters[0] should be identical, but they aren't.")
	}
	// Check how many positions are referenced
	wordCount := 0
	for _, wordIndices := range comp.SimilarClusters[0][0] {
		wordCount += len(wordIndices)
	}
	if wordCount != 115 {
		t.Errorf("The keys for comp.SimilarClusters[0][0] reference %d values, but should reference 115.", wordCount)
	}
}

// def comparator_tests
//   comp, son2, son5 = init_comparator
//   assert("load #{son2.filename} as Text") { son2.class == Text }
//   assert("load #{son5.filename} as Text") { son5.class == Text }
//   assert("initialize comparator") { comp.class == Comparator }
//   assert("set @target_word_clusters") do
//     comp.target_word_clusters.class == Array
//   end
//   assert("set @source_word_clusters") do
//     comp.source_word_clusters.class == Array end
//   assert("ensure @target.word_clusters contains Hashes") do
//     comp.target_word_clusters[0].class == Hash
//   end
//   assert("ensure @source.word_clusters contains Hashes") do
//     comp.source_word_clusters[0].class == Hash
//   end
//   assert("check similar_clusters.class is Array") do
//     comp.similar_clusters.class == Array
//   end
//   assert("check similar_clusters[0].class is Array") do
//     comp.similar_clusters[0].class == Array
//   end
//   assert("check similar_clusters[0][0].class is Hash") do
//     comp.similar_clusters[0][0].class == Hash
//   end
//   assert("avoid erroneous $son2--$son5 match") do
//     comp.similar_clusters[0] != [{"beauty"=>38, "then"=>32, "where"=>40},
//                                  {"beauty"=>53, "lusty"=>49, "where"=>58}]
//   end
// end

// puts "\n"
// puts "Comparator Tests:"
// puts "--------------------------------"
// comparator_tests
