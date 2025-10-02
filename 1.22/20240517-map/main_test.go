package main

import "testing"

func TestEmptyMap(t *testing.T) {

	m := map[int]int{}

	a := m[1]

	if a != 0 {
		t.Error("ゼロ値が得られること")
	}
}

func TestNilMap(t *testing.T) {

	var m map[int]int

	a := m[1]

	if a != 0 {
		t.Error("ゼロ値が得られること")
	}
}
