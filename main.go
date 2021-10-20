package main

import "fmt"

func main() {
	mybill := newBill("Jake's bill")
	mybill.updateTip(10)

	fmt.Println(mybill.format())
}
