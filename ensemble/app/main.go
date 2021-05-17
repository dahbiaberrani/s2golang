package main

import (
	"fmt"
	"ensemble"

)

func main() {
	monEns := ensemble.Ensemble{}
	fmt.Println("Nb d'éléments dans monEns = ", monEns.Card())

	monEns.Ajouter(12)
	monEns.Ajouter(100)
	monEns.Ajouter(42)
	monEns.Ajouter(100)  // Tentative d'ajouter un doublo
	fmt.Println("Nb d'éléments dans monEns = ", monEns.Card())
	fmt.Printf("monEns = %v\n", monEns)

	monEns.Oter(12)
	fmt.Println("Nb d'éléments dans monEns = ", monEns.Card())
	fmt.Printf("monEns = %v\n", monEns)

	tonEns := ensemble.Ensemble{}
	tonEns.Ajouter(11)
	tonEns.Ajouter(42)
	tonEns.Ajouter(33)
	fmt.Printf("Intersection de %s et %s = %s\n", monEns, tonEns, monEns.Inter(tonEns))

	autreEns := ensemble.FromSlice([]int{10, 100, 20, 42, 20})
	fmt.Printf("autreEns = %v\n", autreEns)
}