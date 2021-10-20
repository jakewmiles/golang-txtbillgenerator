package main

import "fmt"

func main() {
	mybill := newBill("Jake's bill")
	fmt.Println(mybill.format())
}
