package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	out := ladderLength("hot", "dog", []string{"hot", "dog", "dot"})
	if out != 3 {
		t.Errorf("got %d, want %d", out, 3)
	}
}
