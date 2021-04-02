package comparator

import (
	"testing"
	"github.com/shushcat/rlat/text"
)

// func TestInitComparator(t *testing.T) {
// 	// If the test calls a failure function such as t.Error or t.Fail,
// 	// then the test has failed.
// 	path := "../sonnets/Sonnet IX.txt"
// 	son9 := InitText(path)
// 	if son9.FileName != path {
// 		t.Errorf("FileName is %s; want '%s'.", son9.FileName, path)
// 	}
// 	if len(son9.WordArray) != 118 {
// 		t.Errorf("WordArray is %d, want 118.", len(son9.WordArray))
// 	}
// }

func TestInitComparator(t *testing.T) {
	son2 := text.InitText(".sonnets/Sonnet II.txt")
	son5 := text.InitText(".sonnets/Sonnet V.txt")
	comp = InitComparator(son2, son5, 3, true, 10, 4, false, "", editDist int) Comparator 
//   comp = Comparator.new(son2, son5)
//   return comp, son2, son5
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
