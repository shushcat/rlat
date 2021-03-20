package options

import (
	"flag"
	"errors"
)

var targetPath, sourcePath, stopPath string
var window, minSharedWords, minWordLen, editDist int
var ordered, stemming bool

func ParseFlags() (string, string, int, int, int, bool, bool, string, int, error) {
	handleError := func(err error) (string, string, int, int, int, bool, bool, string, int, error) {
		return "", "", 10, 3, 4, true, false, "", 0, err
	}
	flag.StringVar(&targetPath, "t", "", "target text file")
	flag.StringVar(&sourcePath, "s", "", "source text file")
	if (targetPath == "") || (sourcePath == "") {
		return handleError(errors.New("Target and source files must be specified."))
	}
	flag.IntVar(&window, "w", 10, "number of words in a window")
	flag.IntVar(&minSharedWords, "msw", 3, "minimum number of shared words in a window")
	flag.IntVar(&minWordLen, "mwl", 4, "minimum length of words included in comparison")
	flag.BoolVar(&ordered, "ord", true, "whether shared words must appear in the same order")
	flag.BoolVar(&stemming, "stem", false, "whether to stem words before commparison")
	flag.StringVar(&stopPath, "stop", "", "text file containing stopwords")
	flag.IntVar(&editDist, "dist", 0, "Damarau-Levenschtein edit distance")
	flag.Parse()
	return targetPath, sourcePath, window, minSharedWords, minWordLen, ordered, stemming, stopPath, editDist, nil
}
