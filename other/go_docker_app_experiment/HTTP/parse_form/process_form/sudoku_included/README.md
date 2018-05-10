[//]: # (Image References)
[image_0]: ../misc/test_sudoku_01.PNG

# Sudoku Solver Form
Parse input string (sudoku puzzle data) and attempt to return the solved puzzle.

## Sample Output
![Sample output from POSTMAN][image_0] 

## Use
- Server
    - run `main.go`
- POST request
    - [POSTMAN](https://www.getpostman.com/)
    - set `body`
        - set `data` to a string of sudoku data
        - i.e. ```1 _ 3 _ _ 6 _ 8 _, _ 5 _ _ 8 _ 1 2 _, 7 _ 9 1 _ 3 _ 5 6, _ 3 _ _ 6 7 _ 9 _, 5 _ 7 8 _ _ _ 3 _, 8 _ 1 _ 3 _ 5 _ 7, _ 4 _ _ 7 8 _ 1 _, 6 _ 8 _ _ 2 _ 4 _, _ 1 2 _ 4 5 _ 7 8```
    - hit [send]


### Calling the sudoku solver with the form data
```golang
// sudokuSolver accepts sudoku board data in the form of a string and returns
// a solved board in the form of a map.
func sudokuSolver(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.Form["data"]) > 0 {
		data := r.Form["data"][0]
		solved, err := solveSudoku(data)
		if err != nil {
			fmt.Fprintf(w, "Unable to solve the given puzzle: %v", solved)
		}
		fmt.Fprintf(w, "Success: The solved board is: %v", solved)
	} else {
		fmt.Fprintln(w, "Nothing to see here - no Sudoku data")
	}
}
```

### Image
```
REPOSITORY                 TAG                 IMAGE ID            CREATED             SIZE
jackburdick/test_sudoku    latest              e42ce79ef5f0        5 seconds ago       9.9MB
```
