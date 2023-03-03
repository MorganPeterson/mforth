package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	forth := NewEval()

	for {
		fmt.Printf("> ")

		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("error reading input: %s\n", err.Error())
			return
		}

		text = strings.TrimSpace(text)
		forth.Eval(strings.Split(text, " "))
	}
}

