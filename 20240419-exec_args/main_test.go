package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestRunInTest(t *testing.T) {
	buf := bytes.NewBuffer(nil)

	wd, err := os.Getwd()
	if err != nil {
		t.Error(err)
		return
	}

	cwd, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Error(err)
		return
	}

	prog := wd + "/prog/prog"
	// Argsの一つ目には、プログラムパスが必要
	args := []string{prog, "-f1=AAA", "-f2=BBB", "CCC"}

	cmd := exec.Cmd{
		Path:   args[0],
		Args:   args,
		Stderr: buf,
		Stdout: buf,
		Dir:    cwd,
	}

	err = cmd.Start()
	if err != nil {
		t.Error(err)
	}

	time.Sleep(100 * time.Millisecond)

	err = cmd.Wait()
	if err != nil {
		t.Error(err)
	}

	time.Sleep(100 * time.Millisecond)

	log.Print(buf.Len())
	log.Print(buf.String())
}
