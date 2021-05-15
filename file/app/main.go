package main

import (
	"file/fileprof"
	"fmt"
)

func main() {
	file := fileprof.FromSlice(1, 10, 100)
	file.Ajouter(1000)
	file.Ajouter(2000)
	file.Ajouter(3000)
	file.Ajouter(4000)

	fmt.Println(file)
	fmt.Println(file.NbElts())
	fmt.Println(file.Tete())
	for !file.EstVide() {
		elt, _ := file.Oter()
		fmt.Println(elt)
		fmt.Println(file.ToSlice())
	}

	if elt, err := file.Oter(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(elt)
		
	}
	fmt.Println(file.ToSlice())
}
