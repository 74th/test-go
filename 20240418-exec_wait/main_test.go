package main

import (
	"log"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestExecInterrupt(t *testing.T) {
	cmd := &exec.Cmd{
		Path: "/bin/sleep",
		Args: []string{"/bin/sleep", "3"},
	}

	err := cmd.Start()
	if err != nil {
		t.Errorf("Start: %v", err)
	}

	timing := 0

	go func() {
		err := cmd.Wait()
		if timing == 0 {
			t.Errorf("即座のWaitが終わらないこと: %v", err)
		}

		log.Printf("Wait: %v", err)

		timing = 2
	}()

	time.Sleep(500 * time.Millisecond)

	timing = 1

	err = cmd.Process.Signal(os.Interrupt)
	if err != nil {
		t.Errorf("Signal: %v", err)
	}

	time.Sleep(500 * time.Millisecond)

	if timing != 2 {
		t.Errorf("Waitが終わること: %d", timing)
	}
}

func TestExecKill(t *testing.T) {
	cmd := &exec.Cmd{
		Path: "/bin/sleep",
		Args: []string{"/bin/sleep", "3"},
	}

	err := cmd.Start()
	if err != nil {
		t.Errorf("Start: %v", err)
	}

	timing := 0

	go func() {
		err := cmd.Wait()
		if timing == 0 {
			t.Errorf("即座のWaitが終わらないこと: %v", err)
		}

		log.Printf("Wait: %v", err)

		timing = 2
	}()

	time.Sleep(500 * time.Millisecond)

	timing = 1

	err = cmd.Process.Kill()
	if err != nil {
		t.Errorf("Signal: %v", err)
	}

	time.Sleep(500 * time.Millisecond)

	if timing != 2 {
		t.Errorf("Waitが終わること: %d", timing)
	}
}
