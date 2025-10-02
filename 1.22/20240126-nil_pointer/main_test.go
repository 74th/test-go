package main

import (
	"fmt"
	"testing"
)

type S struct {
	value string
}

func (s *S) PrintValuePointer() string {
	fmt.Println("  called PrintValuePointer")
	return s.value
}

func (s *S) JustCallPointer() string {
	fmt.Println("  called JustCallPointer")
	return "ok"
}

func (s S) PrintValueStruct() string {
	fmt.Println("  called PrintValueStruct")
	return s.value
}
func (s S) JustCallStruct() string {
	fmt.Println("  called JustCallStruct")
	return "ok"
}

func recoverWrap(name string, f func()) {
	fmt.Println("##", name)
	defer func() {
		e := recover()
		if e != nil {
			fmt.Println("recover: ", name)
		}
	}()
	f()
}

func TestNilPointer(t *testing.T) {

	s := S{value: "Value"}
	fmt.Println("## struct-value-1:", s.PrintValuePointer())
	fmt.Println("## struct-value-2:", s.PrintValueStruct())

	s = S{}
	fmt.Println("## struct-zero-1:", s.PrintValuePointer())
	fmt.Println("## struct-zero-2:", s.PrintValueStruct())

	p := &S{value: "Value"}
	fmt.Println("## pointer-value-1:", p.PrintValuePointer())
	fmt.Println("## pointer-value-2:", p.PrintValueStruct())

	// ここまでは当然ながら問題なし

	p = nil
	recoverWrap("pointer-value-1", func() {
		// ポインタメソッドなら、関数には入れるが、s.valueでpanic()
		fmt.Println("pointer-value-1:", p.PrintValuePointer())
	})
	recoverWrap("pointer-value-2", func() {
		// 構造体メソッドなら、呼び出した時点で、panicする
		fmt.Println("pointer-value-2:", p.PrintValueStruct())
	})

	recoverWrap("pointer-just-call-1", func() {
		// ポインタメソッドなら、呼び出すだけなら、panicしない
		fmt.Println("just-call-1:", p.JustCallPointer())
	})
	recoverWrap("pointer-just-call-2", func() {
		// 構造体メソッドなら、呼び出した時点で、panicする
		fmt.Println("just-call-2:", p.JustCallStruct())
	})
}
