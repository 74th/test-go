package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type Data struct {
	Time time.Time `json:"time",omitempty`
}

type Data2 struct {
	Time *time.Time `json:"time",omitempty`
}

func TestJsonTime(t *testing.T) {
	d := Data{Time: time.Now()}

	b, err := json.Marshal(d)
	fmt.Println("ERR:", err)
	fmt.Println("DATA:", string(b))

	// time.Time は JSON になると、RFC3339 になる
	// DATA: {"t":"2024-01-24T15:54:09.458099807+09:00"}

	d = Data{}
	b, _ = json.Marshal(d)
	fmt.Println("DATA:", string(b))

	// time.Time は omitempty に対応していない
	// DATA: {"t":"0001-01-01T00:00:00Z"}

	now := time.Now()
	d2 := Data2{Time: &now}
	b, _ = json.Marshal(d2)
	fmt.Println("DATA:", string(b))

	d2 = Data2{}
	b, _ = json.Marshal(d2)
	fmt.Println("DATA:", string(b))

	// ポインタにするとnullに対応する
	// DATA: {"t":null}

	d = Data{}
	err = json.Unmarshal(b, &d)
	fmt.Println("ERR:", err)
	fmt.Printf("DATA: %#v\n", d)

	// null はゼロ値になる
	// DATA: main.Data{Time:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}

	d = Data{}
	err = json.Unmarshal([]byte("{}"), &d)
	fmt.Println("ERR:", err)
	fmt.Printf("DATA: %#v\n", d)

	// プロパティがなくてもゼロ値になる
	// DATA: main.Data{Time:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}
}
