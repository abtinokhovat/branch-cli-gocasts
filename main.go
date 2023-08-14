package main

import (
	"branches-cli/cmd"
	"branches-cli/config"
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

	com := flag.String(config.Command.Name, config.Command.Default, config.Command.Description)
	reg := flag.String(config.Region.Name, config.Region.Default, config.Region.Description)
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
