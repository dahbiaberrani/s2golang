package main

import (
	"fmt"
	"pilechainee"
)

func main() {

	// Test de pilechainee
	fmt.Println("Test de pilechainee.Pile...")
	pile2 := pilechainee.FromSlice(1, 10, 100)
	fmt.Printf("Contenu de la pile : %v\n", pile2.ToSlice())
	top, _ := pile2.Top()
	fmt.Printf("La tête de la pile contient %d\n", top)
	pile2.Push(1000)
	fmt.Println("Après pile2.Push(1000) :", pile2.ToSlice())

	fmt.Println("Dépilage de pile2 :")
	for !pile2.IsEmpty() {
		top, _ = pile2.Pop()
		fmt.Println(top)
	}

	if top, err := pile2.Pop(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(top)
	}

	err := pile2.Pop2()
	fmt.Println(err, pile2.ToSlice())
}
