package main

import (
	"encoding/json"
	"testing"
)

type T1 struct {
	A struct {
		B json.RawMessage
	}
}

type T2 struct {
	A json.RawMessage
}

func TestRun(t *testing.T) {
	data := `{"A": {"B": [1,2,3]}}`

	var s1 T1
	err := json.Unmarshal([]byte(data), &s1)
	if err != nil {
		t.Errorf("Failed to marshal JSON: %v", err)
		return
	}
	t.Logf("JSON output: %#v", s1)

	var b1 []int
	err = json.Unmarshal(s1.A.B, &b1)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
		return
	}
	t.Logf("Unmarshaled int slice: %#v", b1)

	var s2 T2
	err = json.Unmarshal([]byte(data), &s2)
	if err != nil {
		t.Errorf("Failed to marshal JSON: %v", err)
		return
	}

	t.Logf("JSON output: %#v", s1)
}
