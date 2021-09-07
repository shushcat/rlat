package distance

import "testing"

func TestDamlev(t *testing.T) {
	var dist int
	dist = DamLev("a", "b")
	if dist != 1 {
		t.Errorf("DamLev('a', 'b') is %d; want 1.", dist)
	}
}
