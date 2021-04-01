
package main

import (
	"fmt"
	"listedbchainee"

) 
func main() {
	liste := Liste{}
	liste.Cons(12)
	liste.Cons(42)
	liste.Cons(55)
	fmt.Println(liste)
	fmt.Printf("Longueur de la liste = %d\n", liste.Length())

	liste.Append(100)
	liste.Append(200)
	fmt.Println(liste)

	liste2 := Liste{}
	liste2.Remove(100)
	fmt.Println("liste2 =", liste2)

	liste3 := FromSlice([]int{10, 20, 3, 1})
	fmt.Println(liste3)

	liste3.Remove(10)
	fmt.Println(liste3)
	liste3.Remove(3)
	fmt.Println(liste3)
	liste3.Remove(42)
	fmt.Println(liste3)

}
