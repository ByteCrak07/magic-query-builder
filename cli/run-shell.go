package cli

import (
	"bufio"
	"fmt"
	"magic-query/generator"
	"os"
)

func RunShell() {
	fmt.Println("Magic Query Builder")
	fmt.Printf("Welcome to Interactive shell. Use Ctrl+C to close the shell or use \"%v --help\" for help\n", os.Args[0])
	fmt.Println("Press enter again after input for output query")

	// create empty slice
	var input []string

	for {
		tempInput := getInput()

		if tempInput == "" {
			output := generator.HandleQuery(input)

			// clear input slice for next query
			input = nil

			deletePreviousLine()

			fmt.Println(output)
		} else {
			input = append(input, tempInput)
		}
	}
}

func deletePreviousLine() {
	// const LINE_UP = "\033[1A"
	// const LINE_CLEAR = "\x1b[2K"
	// fmt.Print(LINE_UP, LINE_CLEAR)
}

func getInput() string {
	fmt.Print(">>> ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	return line
}
