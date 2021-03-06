// Author: Alangi Derick
// Description: Server1 is a minimal "echo" and counter server.

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:6000", nil))
}

// handler echoes the Path component of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock() // lock mutex

	fmt.Fprintf(w, "%s   %s   %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	count++
	mu.Unlock() // unlock mutex
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock() // lock mutex
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
