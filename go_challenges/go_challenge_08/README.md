[//]: # (Image References)
[image_0]: ./misc/solveSudoku_picture_cropped.png


# Go Challenge #8: Sudoku Solver
![before and after][image_0] 
Original Project Information can be found [here](http://golang-challenge.org/go-challenge8/). The main functionality is to solve a Sudoku puzzle from a given input.

## Running the tests
Running `go test` from within the `./sudoku` directory will run the main tests and will output the included sample starting and finishing boards.

## Strategy
The solver `solveSudoku()` attempts to solve the given test board using constraint propagation.
- Cycle 1 `reduce()`:
    1. Produce all possible values for each box
    2. Reduce the possibilities for each box using several strategies
        - `eliminate()`
            - If a given box has been solved, all peers will have the solved value removed from its potential solution options.
        - `onlyChoice()`
            - For any given unit, if only one box has the option to be solved for a given value, then that box will be considered solved with that value.
        - `nakedGroup()` (if unsolved)
            - Implement the [naked twins](http://www.sudokudragon.com/tutorialnakedtwins.htm) strategy. If a given unit has two boxes with only two options, then we know that one box will contain one of each potential solution.  Therefore, any other box in that unit will not be solved with either of these two values.  Naked twins implements this logic and removes those options from the other boxes in a given unit.  This methodology has been extended for "triplets" as well.
- Cycle 2 `search()`: (if still unsolved)
    - Copy board and guess at a solution for a given box (if unsolved)
        1. Select box with fewest possible solution options
        2. Copy the board
        3. Recursively attempt to solve the board with the guessed solution using above 'Cycle 1' 

### Included test dataset
* `puzzle_01.txt` Provided | [link](http://golang-challenge.org/go-challenge8/)
* `puzzle_02.txt` Hard | [link](http://www.websudoku.com/?level=3&set_id=1047193714)
* `puzzle_03.txt` Evil | [link](http://www.websudoku.com/?level=4&set_id=2508589900)
* `puzzle_04.txt` Evil | [link](http://www.websudoku.com/?level=4&set_id=1776784296)


### TODO:
- Improve `var` names and documentation 
- Extend tests for each function
- Pass only the parsed board to the main function
- Stretch
    - Dockerize
    - Build front end
    - Run on the web
- Improve return values to include unsolved case