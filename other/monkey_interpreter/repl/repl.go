package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/JackBurdick/monkey_interpreter/lexer"
	"github.com/JackBurdick/monkey_interpreter/token"
)

const PROMPT = ">> "

// Start reads from the input (until newline), and passes the input to an
// instance of our lexer before printing all the tokens.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
