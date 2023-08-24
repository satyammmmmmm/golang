package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_Main(t *testing.T) {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	main()
	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdout
	if !strings.Contains(output, "34320.00") {
		t.Error("wrong balance returneed")
	}

}

// package main

// import (
// 	"sync"
// 	"testing"

// func Test_updateMessage(t *testing.T) {
// 	msg = "hello,world"
// 	var mutex sync.Mutex
// 	//var wg sync.WaitGroup
// 	wg.Add(2)
// 	go updateMessage("goodbye,cruel world", &mutex)
// 	go updateMessage("goodbye,cruel world", &mutex)
// 	wg.Wait()

// 	if msg != "goodbye,cruel world" {
// 		t.Error("incoorect value in msg ")
// 	}

// }
