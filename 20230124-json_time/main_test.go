package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type D struct {
	T time.Time `json:"t"`
}

func TestJsonTime(t *testing.T) {
	d := D{T: time.Now()}

	b, err := json.Marshal(d)
	fmt.Println("ERR:", err)
	fmt.Println("DATA:", string(b))
	// ERR: <nil>
	// DATA: {"t":"2024-01-24T15:54:09.458099807+09:00"}

	// time.Time は JSON になると、RFC3339 になる
}
