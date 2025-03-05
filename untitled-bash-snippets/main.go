package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"mvdan.cc/sh/v3/syntax"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	reader := bufio.NewReader(file)

	f, err := syntax.NewParser(syntax.KeepComments(true)).Parse(reader, "")

	writer := io.Writer(os.Stdout)
	printer := syntax.NewPrinter()
	syntax.Walk(f, func(node syntax.Node) bool {
		switch node := node.(type) {
		case *syntax.Stmt:
			switch cmd := node.Cmd.(type) {
			case *syntax.FuncDecl:
				if strings.HasPrefix(cmd.Name.Value, "cmd:") {
					if x, ok := cmd.Body.Cmd.(*syntax.Block); ok {
						for _, cmd := range x.Stmts {
							printer.Print(writer, cmd.Cmd)
							if len(cmd.Comments) > 0 {
								fmt.Fprintf(writer, "\t%s\n", strings.TrimSpace(cmd.Comments[0].Text))
							} else {
								fmt.Fprintf(writer, "\n")
							}
						}
					}
				}
			}
		}
		return true
	})
}
