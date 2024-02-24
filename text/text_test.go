package text

import (
	"testing"
	// "reflect"
)

func TestInitText(t *testing.T) {
	// If the test calls a failure function such as t.Error or t.Fail,
	// then the test has failed.
	path := "../txts/sonnets/Sonnet IX.txt"
	son9 := InitText(path)
	if son9.FileName != path {
		t.Errorf("FileName is %s; want '%s'.", son9.FileName, path)
	}
	if len(son9.WordArray) != 118 {
		t.Errorf("WordArray is %d, want 118.", len(son9.WordArray))
	}
}
