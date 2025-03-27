package main

import (
	"testing"
	"time"
)

func testProcessNumber(t *testing.T) {
	expected := 10
	input := 5
	output := make(chan int)

	ProcessNumber(input, output)

	select {
	case result := <-output:
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("Test timed out")
	}
}
