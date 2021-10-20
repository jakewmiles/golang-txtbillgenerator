package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("What's the bill's name?: ", reader)
	b := newBill(name)
	fmt.Println("Created new bill: ", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item) (s - save bill) (t - add tip): ", reader)
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		fmt.Println(name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Item name: ", reader)
		fmt.Println(tip)
		promptOptions(b)
	case "s":
		fmt.Println("s selected")
	default:
		fmt.Printf("%v not an option \n", opt)
		promptOptions(b)
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
}
