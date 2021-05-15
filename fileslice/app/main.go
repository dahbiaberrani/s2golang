package main

import (
	"fileslice"
	"fmt"
)

func main() {
	file := fileslice.FromSlice([] int{1, 10, 100})
	file.Ajouter(1000)

	fmt.Println(file.ToSlice())

	for !file.EstVide() {
		elt, _ := file.Oter()
		fmt.Println(elt)
	}

	if elt, err := file.Oter(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(elt)
	}
}
