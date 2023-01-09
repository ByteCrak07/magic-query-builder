package main

import (
	"bufio"
	"fmt"
	"os"
)

func runShell() {
	fmt.Println("Magic Query Builder")
	fmt.Printf("Welcome to Interactive shell. Use Ctrl+C to close the shell or use \"%v --help\" for help\n", os.Args[0])

	for {
		input := getInput()
		output := generateQuery(input)
		fmt.Println(output)
	}
}

func getInput() string {
	fmt.Print(">>> ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	return line
}

func generateQuery(input string) string {
	output := fmt.Sprintf("{\n\t\"__query__\": %v\n}", input)

	return output
}
