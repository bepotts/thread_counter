package counter

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"thread_counter/counter"
)

// tests the method NonThreadedCount in the counter package
func TestNonThreadedCount(t *testing.T) {
	var buffer bytes.Buffer
	counter.NonThreadedCount(100, &buffer)
	expected := "Done."

	if !strings.Contains(buffer.String(), expected) {
		fmt.Printf("Expect %q, got %q", expected, buffer.String())
	}
}

// tests the method ThreadedCount in the counter package
func TestThreadedCount(t *testing.T) {
	var buffer bytes.Buffer
	counter.ThreadedCount(100, &buffer)
	expected := "Done."

	if !strings.Contains(buffer.String(), expected) {
		fmt.Printf("Expect %q, got %q", expected, buffer.String())
	}
}
