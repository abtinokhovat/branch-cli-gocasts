package cmd

import (
	"branches-cli/internal/branch"
	"branches-cli/internal/region"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Command struct {
	regionService *region.Service
	branchService *branch.Service
	scanner       *bufio.Scanner
}

func NewCommand(regionService *region.Service, branchService *branch.Service, scanner *bufio.Scanner) *Command {
	return &Command{
		regionService: regionService,
		branchService: branchService,
		scanner:       scanner,
	}
}

func (c *Command) ListBranches() {
	branches, err := c.branchService.GetAllBranches()
	if err != nil {
		return
	}
	for _, brn := range branches {
		fmt.Printf("\n===============\n")
		fmt.Println(brn.String())
	}
}
func (c *Command) GetBranch() {
	str := c.scan("Enter the id of the branch:")

	id, err := strconv.Atoi(str)
	if err != nil {
		fmt.Sprintln(err)
	}

	brn, err := c.branchService.GetBranchDetail(id)
	if err != nil {
		fmt.Sprintln(err)
	}

	if brn != nil {
		fmt.Println(brn.String())
	} else {
		fmt.Printf("Branch:#%d not found\n", id)
	}

}
func (c *Command) CreateBranch(region *region.Region) {
	// get a new id
	id, err := c.branchService.NewId()
	if err != nil {
		fmt.Sprintln(err)
	}

	date := time.Now().String()

	// get the name, num of employees
	name := c.scan("Enter the Name of the branch:")
	phone := c.scan("Enter the phone of the branch")
	numOfEmpStr := c.scan("Enter the Number of employees in the branch")

	numOfEmp, err := strconv.Atoi(numOfEmpStr)
	if err != nil {
		fmt.Sprintln(err)
	}

	regionId := region.Id
	if err != nil {
		fmt.Sprintln(err)
	}

	b := branch.New(id, name, phone, date, numOfEmp, regionId)
	err = c.branchService.CreateBranch(b)
	if err != nil {
		fmt.Sprintln(err)
	}

	fmt.Printf("Branch #%d Created: %v", id, b)
}
func (c *Command) EditBranch(region *region.Region) {

}
func (c *Command) GiveStatus(region *region.Region) {
	branches, err := c.branchService.ListBranchesInRegion(region)
	if err != nil {
		fmt.Sprintln(err)
	}

	var numOfBranches int
	if branches != nil {
		numOfBranches = len(branches) + 1
	} else {
		numOfBranches = 0
	}

	var sumNumOfEmp int
	for _, b := range branches {
		sumNumOfEmp += b.NumberOfEmployees
	}

	fmt.Printf(
		"#️⃣%d-%s\n----------------\nNumber of branches: %d, and the number of employees working on them are a total of %d",
		region.Id,
		region.Name,
		numOfBranches,
		sumNumOfEmp,
	)
}
func (c *Command) Execute(command string, regionIdStr string) {
	id, err := strconv.Atoi(regionIdStr)
	if err != nil {
		fmt.Sprintln(err)
	}

	r, err := c.regionService.GetRegionDetail(id)
	if err != nil {
		fmt.Sprintln(err)
	}

	switch command {
	case "list":
		fmt.Printf("\nBranches:\n")
		c.ListBranches()
	case "get":
		fmt.Printf("\nBranch:\n")
		c.GetBranch()
	case "create":
		fmt.Printf("\nCreated:\n")
		c.CreateBranch(r)
	case "edit":
		fmt.Printf("\nEdited:\n")
		c.EditBranch(r)
	case "status":
		fmt.Printf("\nStatus:\n")
		c.GiveStatus(r)
	case "exit":
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Printf("Invalid Command for execution")
	}
}
func (c *Command) scan(prompt string) string {
	fmt.Println(prompt)
	c.scanner.Scan()
	return c.scanner.Text()
}
