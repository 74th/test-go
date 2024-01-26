package main

import (
	"fmt"
	"testing"
)

const (
	typeA = iota
	typeB
)

var (
	typesOnlyA = []int{typeA}
	typesAll   = []int{typeA, typeB}
)

func TestSliceArray(t *testing.T) {

	var slice []int

	slice = typesOnlyA
	fmt.Println("typesOnlyA:", slice)
	slice = typesAll
	fmt.Println("typesOnlyA:", slice)

}
