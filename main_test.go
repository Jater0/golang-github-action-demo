package main

import "testing"

func TestCat(t *testing.T) {
	saying := Cat()
	if saying != "Fuck me"{
		t.Errorf("Cat say %s", saying)
	}
}