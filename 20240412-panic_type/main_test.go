package main

import "testing"

func panicFunc() (e interface{}) {
	defer func() {
		r := recover()
		e = r
	}()

	make([]int, 0)[1] = 1

	return
}

func TestPanicType(t *testing.T) {

	e := panicFunc()

	if e == nil {
		t.Errorf("panicFunc() should panic")
		return
	}

	// 実行時エラーはちゃんとerror型になる
	err, ok := e.(error)
	if !ok {
		t.Errorf("panic should error")
		return
	}

	t.Logf("panic error: %v", err)
}
