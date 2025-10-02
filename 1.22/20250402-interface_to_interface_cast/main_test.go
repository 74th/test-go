package main

import "testing"

type Instance struct {
}

func (a *Instance) DoAAA() {
}

func (a *Instance) DoBBB() {
}

type AAARunner interface {
	DoAAA()
}

type BBBRunner interface {
	DoBBB()
}

type CCCRunner interface {
	DoCCC()
}

func TestRun(t *testing.T) {

	instance := &Instance{}

	aaa := AAARunner(instance)

	_, ok := aaa.(BBBRunner)
	if !ok {
		t.Fatal("変換元のインターフェイスにはなくても、キャストできる")
	}

	_, ok = aaa.(CCCRunner)
	if ok {
		t.Fatal("もちろんエラー")
	}
}
