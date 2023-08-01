package main

import (
	"flag"
	"fmt"
)

func main() {
	command := flag.String("command", "exit", "Command name that can be executed")
	region := flag.String("region", "Tehran", "the city that you want to retrieve the data of branches from")

	flag.Parse()
	fmt.Println("command:", *command)
	fmt.Println("region:", *region)
}
