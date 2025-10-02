package main

import (
	"io"
	"log"
	"sync"
	"testing"
	"time"
)

func TestIOPipe(t *testing.T) {
	reader, writer := io.Pipe()

	go func() {
		writer.Write([]byte("Hello"))
	}()

	buf := make([]byte, 1000)

	n, err := reader.Read(buf)
	if err != nil {
		t.Errorf("何か書いてある時はEOFにならない: %v", err)
	}

	if n != 5 {
		t.Errorf("5 byteであること: %d", n)
	}

	var afterClose bool

	wait := &sync.WaitGroup{}
	wait.Add(1)

	log.Printf("@@1")

	// reader.Close()するまでは、io.EOFが返らず、reader.Read() は待たされる
	go func() {
		defer wait.Done()

		n, err = reader.Read(buf)

		if afterClose {
			if err != io.EOF {
				t.Errorf("CloseするとEOFが返ること: %s", err)
			}
			if n > 0 {
				t.Errorf("EOFが返る時: %s", err)
			}
		} else {
			t.Errorf("空でない場合は、Closeするまで取れないこと: %d, %s", n, err)
		}
	}()

	time.Sleep(time.Millisecond)

	afterClose = true

	// writer.Close() して初めて、readerでEOFが返る
	err = writer.Close()
	if err != nil {
		t.Error(err)
	}

	wait.Wait()
}
