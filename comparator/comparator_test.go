package comparator

import (
	"reflect"
	"testing"
)

func TestDissimilarComparator(t *testing.T) {
	son2 := "../txts/sonnets/Sonnet II.txt"
	son5 := "../txts/sonnets/Sonnet V.txt"
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

func TestSonnet2IdentityComparator(t *testing.T) {
	son2 := "../txts/sonnets/Sonnet II.txt"
	comp := InitComparator(son2, son2, 3, true, 10, 1, false, "", 0)
	// There should be only one cluster
	if len(comp.SimilarClusters) != 1 {
		t.Errorf("len(comp.SimilarClusters) is %d, but should be 1", len(comp.SimilarClusters))
	}
	// Make sure the last word of the sonnet is found
	if comp.SimilarClusters[0][0]["cold"] == nil {
		t.Errorf("The last word of Sonnet II is \"cold\", but comp.SimilarClusters[0][0][\"cold\"] is %d", comp.SimilarClusters[0][0]["cold"])
	}
	// Make sure the cluster's elements are equal
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
	if len(comp.SimilarClusters[0][0].FlatValues()) != 115 || 
		len(comp.SimilarClusters[0][1].FlatValues()) != 115 {
		t.Errorf("There should be 115 keys in FlatValues() for both source and target texts.")
	}
}

func TestSonnet1IdentityComparator(t *testing.T) {
	son1 := "../txts/sonnets/Sonnet I.txt"
	comp := InitComparator(son1, son1, 3, true, 10, 1, false, "", 0)
	// There should be only one cluster
	if len(comp.SimilarClusters) != 1 {
		t.Errorf("len(comp.SimilarClusters) is %d, but should be 1", len(comp.SimilarClusters))
	}
	// Make sure the last word of the sonnet is found
	if comp.SimilarClusters[0][0]["thee"] == nil {
		t.Errorf("The last word of Sonnet I is \"thee\", but comp.SimilarClusters[0][0][\"thee\"] is %d", comp.SimilarClusters[0][0]["thee"])
	}
	// Make sure the cluster's elements are equal
	if !reflect.DeepEqual(comp.SimilarClusters[0][0], comp.SimilarClusters[0][1]) {
		t.Error("The pair at comp.SimilarClusters[0] should be identical, but they aren't.")
	}
	// Check how many positions are referenced
	wordCount := 0
	for _, wordIndices := range comp.SimilarClusters[0][0] {
		wordCount += len(wordIndices)
	}
	if wordCount != 106 {
		t.Errorf("The keys for comp.SimilarClusters[0][0] reference %d values, but should reference 106.", wordCount)
	}
}
