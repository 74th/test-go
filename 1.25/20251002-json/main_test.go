package main

import (
	"encoding/json"
	"testing"
)

// T1を含むT2の形であっても、JSONに変換したら、並列にプロパティが並ぶかどうか
type T1 struct {
	A int
	B string
}

type T2 struct {
	T1
	C float64
}

func TestRun(t *testing.T) {
	data := &T2{
		T1: T1{
			A: 42,
			B: "hello",
		},
		C: 3.14,
	}

	output, err := json.Marshal(data)
	if err != nil {
		t.Errorf("Failed to marshal JSON: %v", err)
		return
	}

	// {"A":42,"B":"hello","C":3.14}
	t.Logf("JSON output: %s", output)
}
