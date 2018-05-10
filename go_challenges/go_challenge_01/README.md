[//]: # (Image References)
[image_0]: ./misc/go_challenge_01_overview.png

## Go Challenge #1: Binary Decoder
![sample decoding output][image_0] 
Original Project Information can be found [here](http://golang-challenge.org/go-challenge1/). The main functionality is to decode a pattern from a given binary backup.

## Running the tests
Running `go test` from within the `./drum` directory will run the main tests.

## Strategy
First, a header is decoded from the binary backup `decodePattern()` -- the header information provides information regarding the body/instrument information to be decoded. The instruments patterns are then decoded with `decodeInstrument()`.  The final decoded information is then formatted according to the specified requirements with the `Stringer` interface.

## Visualizing results
For convenience, a print function `fmt.Println(decoded)` has been placed in `drum_test.go`. If uncommented, an output, similar to the output above, will be displayed in the standard output.

### Included test dataset
* `./drum/fixtures/patern_1.splice` Provided | [link](http://golang-challenge.org/go-challenge1/)
* `./drum/fixtures/patern_2.splice` Provided | [link](http://golang-challenge.org/go-challenge1/)
* `./drum/fixtures/patern_3.splice` Provided | [link](http://golang-challenge.org/go-challenge1/)
* `./drum/fixtures/patern_4.splice` Provided | [link](http://golang-challenge.org/go-challenge1/)
* `./drum/fixtures/patern_5.splice` Provided | [link](http://golang-challenge.org/go-challenge1/)

### Note:
- This solution is considered 'complete' -- there is no planned improvement or functionality expansion
- Future improvements could include;
    - Writing an encoder
        - Reverse the process: pattern -> binary backup
    - Modifying the tracks 
        - A [proposed](http://golang-challenge.org/go-challenge1/) expansion is to 'add more cowbell'