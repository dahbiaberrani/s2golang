
package main

import (
	"fmt"
	"listedbchainee"

) 


func main() {
	liste := listedbchainee.Liste{}
	fmt.Println(liste)
	liste.Cons(15)
	liste.Cons(3)
	liste.Cons(12)
	liste.Cons(2)
	fmt.Println(liste)
	fmt.Printf("Longueur de la liste = %d\n", liste.Length())



    liste1 := listedbchainee.Liste{}
	liste1.Remove(2)
	fmt.Println("liste1 =", liste1)
	liste.Append(100)
	liste.Append(200)
	fmt.Println("liste =" ,liste)

	liste.Remove(2)
	fmt.Println("liste =", liste)
	liste.Cons(2)
	liste.Remove(3)
	fmt.Println("liste =", liste)
	liste.Remove(15)
	fmt.Println("liste =", liste)

	fmt.Printf("Longueur de la liste = %d\n", liste.Length())


	liste3 := listedbchainee.FromSlice([]int{10, 20, 3, 1})
	fmt.Println(liste3)

}

