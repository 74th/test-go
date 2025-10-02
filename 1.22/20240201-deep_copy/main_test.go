package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

type S struct {
	Value       int
	Pointer     *int
	ZeroPointer *int
	NilPointer  *int
}

func TestGobDeepcopy(t *testing.T) {
	n := 10
	zero := 0
	s0 := S{Value: 1, Pointer: &n, ZeroPointer: &zero, NilPointer: nil}

	buf := bytes.NewBuffer(nil)
	err := gob.NewEncoder(buf).Encode(s0)
	if err != nil {
		panic(err)
	}

	fmt.Println("BUFLEN:", buf.Len())

	var s1 S

	err = gob.NewDecoder(buf).Decode(&s1)
	if err != nil {
		panic(err)
	}

	fmt.Println("Pointer:", s0.Pointer, s1.Pointer)             // ポインターはコピーされる
	fmt.Println("NilPointer:", s0.NilPointer, s1.NilPointer)    // nilはnil
	fmt.Println("ZeroPointer:", s0.ZeroPointer, s1.ZeroPointer) // ゼロ値のポインターはnilになる
}

func TestBsonDeepcopy(t *testing.T) {
	n := 10
	zero := 0
	s0 := S{Value: 1, Pointer: &n, ZeroPointer: &zero, NilPointer: nil}

	b, err := bson.Marshal(s0)
	if err != nil {
		panic(err)
	}

	fmt.Println("BUFLEN:", len(b))

	var s1 S

	err = bson.Unmarshal(b, &s1)
	if err != nil {
		panic(err)
	}

	fmt.Println("Pointer:", s0.Pointer, s1.Pointer)             // ポインターはコピーされる
	fmt.Println("NilPointer:", s0.NilPointer, s1.NilPointer)    // nilはnil
	fmt.Println("ZeroPointer:", s0.ZeroPointer, s1.ZeroPointer) // ゼロ値のポインターはnilになる
}
