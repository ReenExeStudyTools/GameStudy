package player

import "testing"

func TestPlayer(t *testing.T) {
	if player() != [4]int{5, 5, 5, 1 << 31} {
		t.Fatal("fail")
	}
}