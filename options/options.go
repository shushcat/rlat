package options

import (
	"flag"
)

var targetPath, sourcePath, stopPath string
var window, minSharedWords, minWordLen, editDist int
var ordered, stemming bool

// TODO Add error handling to require source and target file specification.
func ParseFlags() (string, string, int, int, int, bool, bool, string, int) {
	flag.StringVar(&targetPath, "t", "", "target text file")
	flag.StringVar(&sourcePath, "s", "", "source text file")
	flag.IntVar(&window, "w", 10, "number of words in a window")
	flag.IntVar(&minSharedWords, "msw", 3, "minimum number of shared words in a window")
	flag.IntVar(&minWordLen, "mwl", 4, "minimum length of words included in comparison")
	flag.BoolVar(&ordered, "ord", true, "whether shared words must appear in the same order")
	flag.BoolVar(&stemming, "stem", false, "whether to stem words before commparison")
	flag.StringVar(&stopPath, "stop", "", "text file containing stopwords")
	flag.IntVar(&editDist, "dist", 0, "Damarau-Levenschtein edit distance")
	flag.Parse()
	return targetPath, sourcePath, window, minSharedWords, minWordLen, ordered, stemming, stopPath, editDist
}
