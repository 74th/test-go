package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type Data struct {
	Time time.Time `json:"t",omitempty`
}

type Data2 struct {
	Time *time.Time `json:"t",omitempty`
}

func TestJsonTime(t *testing.T) {
	d := Data{Time: time.Now()}

	b, err := json.Marshal(d)
	fmt.Println("ERR:", err)
	fmt.Println("DATA:", string(b))
	// ERR: <nil>
	// DATA: {"t":"2024-01-24T15:54:09.458099807+09:00"}

	// time.Time は JSON になると、RFC3339 になる

	d = Data{}
	b, _ = json.Marshal(d)
	fmt.Println("DATA:", string(b))

	// DATA: {"t":"0001-01-01T00:00:00Z"}

	// time.Time は omitempty に対応していない

	now := time.Now()
	d2 := Data2{Time: &now}
	b, _ = json.Marshal(d2)
	fmt.Println("DATA:", string(b))

}
