package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
)

const PROMPT = "(~˘▾˘)~ >> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		// this is the AST
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		// evaluated AST
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, "\n")
			io.WriteString(out, "Evaluated result: \n")
			io.WriteString(out, "\n")
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

		io.WriteString(out, "\n")
		io.WriteString(out, "AST: \n")
		io.WriteString(out, "\n")
		io.WriteString(out, program.String())
		io.WriteString(out, "\n\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "SOMETHING BAD HAPPENED!!! \n")
	io.WriteString(out, "parser errors: \n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
