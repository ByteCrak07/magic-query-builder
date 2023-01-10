package cli

import (
	"fmt"
	"magic-query/generator"
	"os"
	"strings"
)

func HandleArguments() {
	// stop execution of fn if no arguments are provided
	if len(os.Args) == 1 {
		return
	}

	arg := os.Args[1]

	// if a flag is provided as arguments else argument is a filename
	if strings.HasPrefix(arg, "-") {
		if arg == "--help" {
			printHelp()
		} else if arg == "--serve" {
			var port string
			if len(os.Args) == 3 {
				port = os.Args[2]
			} else {
				port = "5000"
			}

			handleServer(port)
		} else {
			printArgError()
		}
	} else {
		fileBytes, err := os.ReadFile(arg)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// split each line in the file to string slices
		sliceData := strings.Split(string(fileBytes), "\n")
		output := generator.HandleQuery(sliceData)

		fmt.Println(output)
		os.Exit(0)
	}

}

func printHelp() {
	fmt.Printf("usage: %v [option]\n", os.Args[0])
	fmt.Println("Available options")
	fmt.Printf("-       : simply run \"%v\" to run program in an interactive shell\n", os.Args[0])
	fmt.Println("<file>  : read and run program from an input file")
	fmt.Println("--serve : start a server at 5000 (default port) or at a given port no (use: --serve <port>)")
	fmt.Println("--help  : show help")
	os.Exit(0)
}

func printArgError() {
	fmt.Printf("Unknown option: %v\n", os.Args[1])
	fmt.Printf("Try \"%v --help\" for help\n", os.Args[0])
	os.Exit(1)
}
