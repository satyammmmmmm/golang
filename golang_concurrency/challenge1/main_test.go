package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func Test_UpdateMessage(t *testing.T) {
	wg.Add(1)
	go updateMessage("epsilon", &wg)
	wg.Wait()
	if msg != "epsilon" {
		t.Error("Expected to find epsilon, but it is not there")
	}

}
func Test_PrintMessage(t *testing.T) {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	msg = "sps"
	printMessage()
	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdout
	if !strings.Contains(output, "sps") {
		t.Error("not working correctly")
	}
}
func Test_Main(t *testing.T) {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	main()
	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdout

	if !strings.Contains(output, "Hello, universe!") {
		t.Error("Expected to find Hello, universe!, but it is not there")
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Error("Expected to find Hello, cosmos!, but it is not there")
	}

	if !strings.Contains(output, "Hello, world!") {
		t.Error("Expected to find Hello, world!, but it is not there")
	}
}
