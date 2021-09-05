package options

import (
	"testing"
)

func TestDefaultOptions(t *testing.T) {
	targetPath, sourcePath, window, minSharedWords, minWordLen, ordered, stemming, stopPath, editDist, err := ParseFlags()
	if targetPath != "" {
		t.Errorf("targetPath is %s; should be empty.", targetPath)
	}
	if sourcePath != "" {
		t.Errorf("sourcePath is %s; should be empty.", sourcePath)
	}
	if window != 10 {
		t.Errorf("window is %d; want %d.", window, 10)
	}
	if minSharedWords != 3 {
		t.Errorf("minSharedWords is %d; want '%d'.", minSharedWords, 3)
	}
	if minWordLen != 4 {
		t.Errorf("minWordLen is %d; want '%d'.", minWordLen, 4)
	}
	if ordered != true {
		t.Errorf("ordered is %t; want '%t'.", ordered, true)
	}
	if stemming != false {
		t.Errorf("stemming is %t; want '%t'.", stemming, false)
	}
	if stopPath != "" {
		t.Errorf("stopPath is %s; should be empty.", stopPath)
	}
	if editDist != 0 {
		t.Errorf("editDist is %d; want '%d'.", editDist, '0')
	}
	if err == nil {
		t.Error("err is nil but should exist")
	}
}
