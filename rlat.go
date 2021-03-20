package main

import (
	"github.com/shushcat/rlat/comparator"
	"github.com/shushcat/rlat/options"
	"github.com/shushcat/rlat/reports"
	"fmt"
	// "github.com/shushcat/rlat/damlev"
	// "github.com/shushcat/rlat/stemmer"
)

func main() {
	targetPath, sourcePath, window, minSharedWords, minWordLen, ordered, stemming, stopPath, editDist, err := options.ParseFlags()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// fmt.Println(targetPath, sourcePath, window, minSharedWords, minWordLen, ordered, stemming, stopPath, editDist)
	comparator.InitComparator(targetPath, sourcePath, minSharedWords, ordered, window, minWordLen, stemming, stopPath, editDist)
	c := comparator.InitComparator(targetPath, sourcePath, minSharedWords, ordered, window, minWordLen, stemming, stopPath, editDist)
	reports.PrintReport(c)
}
