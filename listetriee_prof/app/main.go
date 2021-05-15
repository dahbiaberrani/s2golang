package main

import (
	"fmt"
	"listetriee_prof"
)

func main() {
	var maListeTriee listetriee_prof.ListeTriee
	maListeTriee.Delete(10)

	maListeTriee = listetriee_prof.FromList([]int{12, 1, 78, 42, 15})

	fmt.Printf("maListeTriée = %v\n", maListeTriee)
	fmt.Println(maListeTriee.Length())

	maListe := maListeTriee.ToList()
	fmt.Printf("maListe = %v\n", maListe)
	fmt.Println(len(maListe))

	maListeTriee.Add(54)
	fmt.Printf("maListeTriée = %v\n", maListeTriee)
	fmt.Println(maListeTriee.Length())

	var val int
	fmt.Print("Entrez une valeur : ")
	fmt.Scan(&val)

	cell := maListeTriee.First(val)
	if  cell != nil {
		fmt.Printf("%v est dans la liste\n", val)
		fmt.Printf("%v est dans la liste\n", cell)
	} else {
		fmt.Printf("%d n'est pas dans la liste\n", val)
	}

	maListeTriee.Replace(100, 0)
	fmt.Println(maListeTriee)
	maListeTriee.Replace(1, 13)
	fmt.Println(maListeTriee)
	maListeTriee.Replace(42, 1)
	fmt.Println(maListeTriee)

	maListeTriee.Delete(100)
	fmt.Println(maListeTriee)
	maListeTriee.Delete(15)
	fmt.Println(maListeTriee)
}
