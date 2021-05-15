package main
import (
"fmt"
"pile/pilechainee"

)

func main() {
	mapile := pilechainee.CreerPile()
	//fmt.Println(mapile.Sommet())
	fmt.Println(mapile)
	fmt.Println(mapile.EstVide())
	mapile.Empiler(55)
	fmt.Println(mapile)
	fmt.Println(mapile.EstVide())
	mapile.Empiler(108)
	fmt.Println(mapile)
	fmt.Println(mapile.ToSlice())
	fmt.Println(mapile.Sommet())
	fmt.Println(mapile.Depiler())
	fmt.Println(mapile)
	fmt.Println(mapile.Depiler())
	fmt.Println(mapile)
	//fmt.Println(mapile.Depiler())
	
	pile := pilechainee.FromSlice([]int{1, 10, 100})
	pile.Empiler(1000)
	fmt.Println(pile)
	fmt.Println(pile.ToSlice()) // [1 10 100 1000]
	for !pile.EstVide() {
		top, _ := pile.Depiler()
		fmt.Println(top) // Affiche de 1000 à 1
	}
	top, err := pile.Depiler();
	if err != nil {
		fmt.Println(err)
	} else { // Affiche "Pile vide"
		fmt.Println(top)
	}
	
}
