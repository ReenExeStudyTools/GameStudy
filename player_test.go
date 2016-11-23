package player

import "testing"

func TestPlayer(t *testing.T) {
	if player() != [1]int{5} {
		t.Fatal("fail")
	}
}