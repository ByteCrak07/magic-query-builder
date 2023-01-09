package main

import "magic-query/cli"

func main() {
	// fn to handle arguments provided while executing the program
	cli.HandleArguments()

	// if no arguments are provided run interacive shell
	cli.RunShell()
}
