package main

import (
	"testing"
	"time"
)

func TestChanString(t *testing.T) {

	pipe1 := make(chan string)
	pipe2 := make(chan string)

	var afterSend bool

	go func() {
		time.Sleep(time.Millisecond)
		afterSend = true
		pipe2 <- "Hello"
	}()

	var s string
	select {
	case s = <-pipe1:
	case s = <-pipe2:
	}
	if !afterSend {
		t.Error("pipeに値を投げるまですすまないこと")
	}
	if len(s) == 0 {
		t.Error("取得できること")
	}
}
