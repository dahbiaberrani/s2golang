package main
import (
"fmt"
"pile/pileprof"
)

func main() {
	mapile := pileprof.CreerPile()
	//fmt.Println(mapile.Sommet())
	fmt.Println(mapile)
	fmt.Println(mapile.IsEmpty())
	mapile.Push(55)
	fmt.Println(mapile)
	fmt.Println(mapile.IsEmpty())
	mapile.Push(108)
	fmt.Println(mapile)
	fmt.Println(mapile.ToSlice())
	fmt.Println(mapile.Top())
	fmt.Println(mapile.Pop())
	fmt.Println(mapile)
	fmt.Println(mapile.Pop())
	fmt.Println(mapile)
	//fmt.Println(mapile.Depiler())
	
	pile := pileprof.FromSlice(1, 10, 100)
	pile.Push(1000)
	fmt.Println(pile)
	fmt.Println(pile.ToSlice()) // [1 10 100 1000]
	for !pile.IsEmpty() {
		top, _ := pile.Pop()
		fmt.Println(top) // Affiche de 1000 à 1
	}
	top, err := pile.Pop();
	if err != nil {
		fmt.Println(err)
	} else { // Affiche "Pile vide"
		fmt.Println(top)
	}
	
}
