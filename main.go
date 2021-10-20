package main

import "fmt"

func main() {
	mybill := newBill("Jake's bill")
	mybill.addItem("french onion", 4.00)
	mybill.addItem("sirloin steak", 16.00)
	mybill.addItem("creme brulee", 7.00)
	mybill.updateTip(10)

	fmt.Println(mybill.format())
}
