package main

import (
	"fmt"
	"listetriee"
)

func main() {
	/*
	var val int
	fmt.Print("Entrez une valeur : ")
	fmt.Scan(&val)                      // On entre 42

	if maListeTriee.First(val) != nil {
		fmt.Printf("%d est dans la liste\n", val)   // 42 est dans la liste
	} else {
		fmt.Printf("%d n'est pas dans la liste\n", val)
	}

	
	fmt.Println(maListeTriee)              // 1 12 15 42 54 78
	maListeTriee.Replace(1, 13)
	fmt.Println(maListeTriee)              // 12 13 15 42 54 78
	maListeTriee.Replace(42, 1)
	fmt.Println(maListeTriee)              // 1 12 13 15 54 78
	*/

	ma_liste :=  ListeTriee{}
	ma_liste.head = nil
	ma_liste.tail = nil


	ma_liste.Add(45)
	ma_liste.Add(40)
	ma_liste.Add(50)
	ma_liste.Add(55)
	ma_liste.Add(65)
	ma_liste.Add(20)
	ma_liste.Add(41)
	ma_liste.Add(62)
	ma_liste.Add(23)
	ma_liste.Add(5)

	fmt.Println(ma_liste)
	maListeTriee := ListeTriee{head:nil,tail:nil}
	maListeTriee.FromList([]int{12, 1, 78, 42, 15})
	fmt.Println(maListeTriee)
	fmt.Println(maListeTriee.Length())     // 5
	maListe := maListeTriee.ToList()
	fmt.Println(maListe)                   // [1 12 15 42 78]



	fmt.Println("=============================")
	maListe1 := ListeTriee{head:nil,tail:nil}
	fmt.Println(maListe1)
	maListe1.Delete(1)
	fmt.Println(maListe1)

	maListe1.Add(12)
	fmt.Println(maListe1)
	maListe1.Delete(1)
	fmt.Println(maListe1)
	maListe1.Delete(12)
	fmt.Println(maListe1)


	maListe1.FromList([]int{12, 1, 78, 42, 15})
	fmt.Println(maListe1)
	maListe1.Delete(1)
	fmt.Println(maListe1)
	maListe1.Delete(15)
	fmt.Println(maListe1)
	maListe1.Delete(78)
	fmt.Println(maListe1)
	maListe1.Add(12)
	maListe1.Add(12)
	maListe1.Add(12)
	fmt.Println(maListe1)

	maListe1.Replace(12, 13)
	fmt.Println(maListe1)


	
}