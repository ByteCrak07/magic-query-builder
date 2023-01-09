package cli

import (
	"fmt"
	"os"
	"strings"
)

func HandleArguments() {
	if len(os.Args) == 1 {
		return
	}

	arg := os.Args[1]

	if strings.HasPrefix(arg, "-") {
		if arg == "--help" {
			printHelp()
		} else {
			printArgError()
		}
	} else {
		data, err := os.ReadFile(arg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// TODO: read from file
		output := generateQuery(string(data))
		fmt.Println(output)
		os.Exit(0)
	}

}

func printHelp() {
	fmt.Printf("usage: %v [option]\n", os.Args[0])
	fmt.Println("Available options")
	fmt.Printf("-      : simply run \"%v\" to run program in an interactive shell\n", os.Args[0])
	fmt.Println("<file> : read and run program from an input file")
	fmt.Println("--help : show help")
	os.Exit(0)
}

func printArgError() {
	fmt.Printf("Unknown option: %v\n", os.Args[1])
	fmt.Printf("Try \"%v --help\" for help\n", os.Args[0])
	os.Exit(1)
}
