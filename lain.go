package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print("lain> ")
		code, _ := reader.ReadString('\n')

		fmt.Println("Executing...")

		fmt.Println("Generating syntax tree")
		parseTree, parseError := ParseStatement(code)
		fmt.Printf("Error: %d\n", parseError)
		fmt.Printf("Identifier: %s\n", GetIdentifier(parseTree.declaration.identifier))

		fmt.Println("Executing code...")
		DummyStatement(parseTree)
	}
}
