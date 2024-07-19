package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Hoge struct {
	PascalCase string
	Snake_Case string
	WithMeta   string `json:"with_meta"`
}

func TestHoge(t *testing.T) {

	h := Hoge{
		PascalCase: "pascal",
		Snake_Case: "snake",
		WithMeta:   "with_meta",
	}

	b, err := json.Marshal(h)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", b)
	// プロパティ名がそのまま出力される
	// {"PascalCase":"pascal","Snake_Case":"snake","with_meta":"with_meta"}
}
