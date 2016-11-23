package player

import "testing"

func TestFibFunc(t *testing.T) {
	if player() != 1 {
		t.Fatalf("fail")
	}
}