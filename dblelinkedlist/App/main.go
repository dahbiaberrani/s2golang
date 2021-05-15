package main

import (
	"dblelinkedlist"
	"fmt"
)

func main() {
	maListe := dblelinkedlist.FromList([]int{12, 3, 42, 2})
	fmt.Println(maListe.Length())    // 4
	fmt.Println(maListe)
	maListe.Append(100)
	maListe.Cons(-1000)
	fmt.Println(maListe.Length())    // 6
	fmt.Println(maListe)

	maTranche := maListe.ToList()
	fmt.Println(maTranche)

	maListe.Delete(-1000)
	maListe.Delete(100)
	maListe.Delete(3)
	fmt.Println(maListe.Length())
	fmt.Println(maListe)

	maListe.Append(666)
	fmt.Println(maListe)

	// test avec une liste d'un seul élément (tete == queue)
	maListe2 := dblelinkedlist.FromList([]int{12})
	maListe2.Delete(12)
	fmt.Println(maListe2.Length())
	maListe2.Append(66)
	fmt.Println(maListe2.Length())
	fmt.Println(maListe2)
	maListe2.Cons(55)
	fmt.Println(maListe2)
	maListe2.Delete(88)
	fmt.Println(maListe2.Length())
	fmt.Println(maListe2)



}
