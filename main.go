package main

import (
	"branches-cli/cmd"
	"branches-cli/internal/branch"
	"branches-cli/internal/region"
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println(`
	
██████╗██████╗ █████╗███╗   ██╗████████╗  ██╗     ████████╗    ██╗
██╔══████╔══████╔══██████╗  ████╔════██║  ██║    ██╔════██║    ██║
██████╔██████╔█████████╔██╗ ████║    ██████████████║    ██║    ██║
██╔══████╔══████╔══████║╚██╗████║    ██╔══██╚════██║    ██║    ██║
██████╔██║  ████║  ████║ ╚████╚████████║  ██║    ╚███████████████║
╚═════╝╚═╝  ╚═╚═╝  ╚═╚═╝  ╚═══╝╚═════╚═╝  ╚═╝     ╚═════╚══════╚═╝
	`)

	fmt.Println("Written by ABTIN OKHOVAT")

	com := flag.String("command", "status", "Command name that can be executed")
	reg := flag.String("reg", "Tehran", "the city that you want to retrieve the data of branches from")
	flag.Parse()

	// setup
	scanner := bufio.NewScanner(os.Stdin)
	branchService := branch.BuildService()
	regionService := region.BuildService()

	command := cmd.NewCommand(regionService, branchService, scanner)

	for {
		command.Execute(*com, *reg)

		fmt.Println("\n\nenter another command: ")
		scanner.Scan()
		if scanner.Text() != "change" {
			*com = scanner.Text()
		} else {
			*reg = scanner.Text()
		}
	}
}
