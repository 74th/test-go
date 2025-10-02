package main

import (
	"log"
	"slices"
	"sort"
	"testing"
)

type Record struct {
	ID   int
	Name string
}

func TestSlices(t *testing.T) {
	rs := []Record{
		{ID: 20, Name: "B"},
		{ID: 10, Name: "A"},
		{ID: 30, Name: "C"},
	}

	slices.SortFunc(rs, func(a, b Record) int {
		return a.ID - b.ID
	})

	log.Printf("slices.SortFunc() rs: %v", rs)

	target := 20
	n, ok := slices.BinarySearchFunc(rs, Record{ID: target}, func(a, b Record) int {
		return a.ID - b.ID
	})
	log.Printf("slices.BinarySearchFunc() target: %d, n: %d, ok: %v", target, n, ok)

	target = 15
	n, ok = slices.BinarySearchFunc(rs, Record{ID: target}, func(a, b Record) int {
		return a.ID - b.ID
	})
	log.Printf("slices.BinarySearchFunc() target: %d, n: %d, ok: %v", target, n, ok)
}

func TestSort(t *testing.T) {
	rs := []Record{
		{ID: 20, Name: "B"},
		{ID: 10, Name: "A"},
		{ID: 30, Name: "C"},
	}

	sort.Slice(rs, func(i, j int) bool {
		return rs[i].ID < rs[j].ID
	})

	log.Printf("sort.Find() rs: %v", rs)

	target := 20
	n, ok := sort.Find(len(rs), func(i int) int {
		return target - rs[i].ID
	})
	log.Printf("sort.Find() target: %d, n: %d, ok: %v", target, n, ok)

	target = 15
	n, ok = sort.Find(len(rs), func(i int) int {
		return rs[i].ID - target
	})
	log.Printf("sort.Find() target: %d, n: %d, ok: %v", target, n, ok)
}

func TestString(t *testing.T) {
	rs := []string{"30", "10", "20"}

	sort.Strings(rs)

	target := "20"
	n := sort.SearchStrings(rs, target)
	log.Printf("sort.SearchString target: %s, n: %d, rs[n]:%s ok: %v", target, n, rs[n], rs[n] == target)

	target = "15"
	n = sort.SearchStrings(rs, target)
	log.Printf("sort.SearchString target: %s, n: %d, rs[n]:%s ok: %v", target, n, rs[n], rs[n] == target)

	// n < len(n) も評価しないとpanicになる
	target = "5"
	n = sort.SearchStrings(rs, target)
	log.Printf("sort.SearchString target: %s, n: %d n<len(rs): %v", target, n, n < len(rs))

	target = "35"
	n = sort.SearchStrings(rs, target)
	log.Printf("sort.SearchString target: %s, n: %d n<len(rs): %v", target, n, n < len(rs))
}
