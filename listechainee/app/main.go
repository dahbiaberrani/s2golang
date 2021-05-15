package main

import (
	"fmt"
	"listechainee"
)



func main() {
	liste := listechainee.Liste{}
	liste.Cons(12)
	liste.Cons(42)
	liste.Cons(55)
	fmt.Println(liste)
	fmt.Printf("Longueur de la liste = %d\n", liste.Length())

	liste.Append(100)
	liste.Append(200)
	fmt.Println(liste)
	fmt.Printf("Longueur de la liste = %d\n", liste.Length())
	liste2 := listechainee.Liste{}
	fmt.Println("longueur de liste2 =", liste2.Length())

	liste3 := listechainee.FromSlice([]int{10, 20, 3, 1})
	fmt.Println(liste3)

	liste3.Remove(10)
	fmt.Println(liste3)
	liste3.Remove(3)
	fmt.Println(liste3)
	liste3.Remove(42)
	fmt.Println(liste3)
	fmt.Printf("Longueur de la liste = %d\n", liste3.Length())

}
