package tests

import "testing"

func TestRandom(t *testing.T) {
	expect := 10
	got := 10

	if got != expect {
		t.Errorf("Expected %d, but got %d\n", expect, got)
	}
}