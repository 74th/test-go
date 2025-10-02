package main

import (
	"math/rand/v2"
	"sort"
	"testing"
	"time"
)

type data struct {
	Key   int64
	Value int64
}

func sortByKeysWithMap(keys []int64, values []data) []data {
	if len(keys) != len(values) {
		return nil
	}

	m := make(map[int64]data, len(keys))
	for i, key := range keys {
		m[key] = values[i]
	}

	result := make([]data, 0, len(keys))
	for _, key := range keys {
		if value, ok := m[key]; ok {
			result = append(result, value)
		}
	}

	return result
}

func sortByKeysWithSortedSlice(keys []int64, values []data) []data {
	if len(keys) != len(values) {
		return nil
	}

	result := make([]data, len(keys))

	sort.Slice(values, func(i, j int) bool {
		return values[i].Key < values[j].Key
	})

	for i, key := range keys {
		n, _ := sort.Find(len(values), func(j int) int {
			return int(key - values[j].Key)
		})
		result[i] = values[n]
	}

	return result
}

func TestRun(t *testing.T) {
	r := rand.New(rand.NewPCG(0, 0))

	keys := make([]int64, 0, 10)
	values := make([]data, 0, 10)
	for i := uint64(0); i < 100; i++ {
		key := r.Int64()

		keys = append(keys, key)
		values = append(values, data{
			Key:   key,
			Value: r.Int64(),
		})
	}

	r.Shuffle(len(values), func(i, j int) {
		values[i], values[j] = values[j], values[i]
	})

	start := time.Now()
	for i := 0; i < 10000; i++ {
		_ = sortByKeysWithMap(keys, values)
	}
	elapsed := time.Since(start)
	t.Logf("sortByKeysWithMap took %s", elapsed)

	start = time.Now()
	for i := 0; i < 10000; i++ {
		_ = sortByKeysWithSortedSlice(keys, values)
	}
	elapsed = time.Since(start)
	t.Logf("sortByKeysWithSortedSlice took %s", elapsed)
}
