package observer

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestObserver(t *testing.T) {
	// This is a hack to capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Make our subject and setup observers
	subject := NewSubject()
	NewBinaryObserver(subject)
	NewHexObserver(subject)

	// Set the state
	subject.SetState(27)

	// More hack to capture stdout
	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old
	out := <-outC

	if !strings.Contains(out, "11011") {
		t.Errorf("The binary observer did not work")
	}

	if !strings.Contains(out, "1b") {
		t.Errorf("The hex observer did not work")
	}
}
