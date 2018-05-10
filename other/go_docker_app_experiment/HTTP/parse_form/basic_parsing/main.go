package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)
	mux.HandleFunc("/form", form)

	http.ListenAndServe(":8080", mux)
}

// handler is the main handler and returns the current time.
func handler(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now().Format("02.01.2006 15:04:05")
	fmt.Fprintf(w, "%s", curTime)
}

// form performs basic form processing.
func form(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for key, value := range r.Form {
		fmt.Printf("%s = %s\n", key, value)
	}
	if len(r.Form["name"]) > 0 {
		fmt.Fprintf(w, "Hi %v!", r.Form["name"][0])
	} else {
		fmt.Fprintln(w, "Why don't you have a name?")
	}
}
