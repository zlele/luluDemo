package mhttp

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

var counter int
var mutex = &sync.Mutex{}
var staticDirectory string

// Replies to the request with the contents of the named file or directory
func serveFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, staticDirectory)
}

func getGoroutineID() uint64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := buf[:n]
	idStr := string(idField)
	var id uint64
	for i := 10; i < len(idStr); i++ {
		if idStr[i] == ' ' {
			break
		}
		id = id*10 + uint64(idStr[i]-'0')
	}
	return id
}

// Just replies with the word 'Hi' as the body of the page
func hiString(w http.ResponseWriter, r *http.Request) {
	goroutineID := getGoroutineID()
	fmt.Fprintf(w, "Hi")
	fmt.Fprintf(w, "Processing request: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Request processed by goroutine %d", goroutineID)
	time.Sleep(10 * time.Second)
}
