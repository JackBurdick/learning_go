package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)
	mux.HandleFunc("/sudokuSolver", sudokuSolver)

	http.ListenAndServe(":8080", mux)
}

// handler is the main handler and returns the current time.
func handler(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now().Format("02.01.2006 15:04:05")
	fmt.Fprintf(w, "%s", curTime)
}

// sudokuSolver accepts sudoku board data in the form of a string and returns
// a solved board in the form of a map.
func sudokuSolver(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.Form["data"]) > 0 {
		data := r.Form["data"][0]

		// TODO: Send this through a load balancer to call container w/compiled function?
		solved, err := SolveSudoku(data)
		if err != nil {
			fmt.Fprintf(w, "Unable to solve the given puzzle: %v", solved)
		}
		fmt.Fprintf(w, "Success: The solved board is: %v", solved)
	} else {
		fmt.Fprintln(w, "Nothing to see here - no Sudoku data")
	}
}
