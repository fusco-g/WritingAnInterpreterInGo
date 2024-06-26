package repl

import (
	"bufio"
	"fmt"
	"interpeter/lexer"
	"interpeter/token"
	"io"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf("Source Code: \n")
		if scanned := scanner.Scan(); !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.NewLexer(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}

}
