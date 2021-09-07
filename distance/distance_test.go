package damlev

import "testing"

func TestDamlev(t *testing.T) {
	var dist int
	dist = Distance("a", "b")
	if dist != 1 {
		t.Errorf("Distance('a', 'b') is %d; want 1.", dist)
	}
}
