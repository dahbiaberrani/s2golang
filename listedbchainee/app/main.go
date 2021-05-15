
package main

import (
	"fmt"
	"listedbchainee"

) 


func main() {

// test la fonction remove avec cas la liste est vide :

    liste1 := listedbchainee.Liste{}
	liste1.Remove(2)
	fmt.Println("liste1 =", liste1)



// test la fonction cons qui ajoute elm
 	liste := listedbchainee.Liste{}
	fmt.Println(liste)
	liste.Cons(15)
	liste.Cons(3)
	liste.Cons(12)
	liste.Cons(2)
	fmt.Println(liste)
	fmt.Println(liste.ReverseString())
// test la longueur de la liste 
	fmt.Printf("Longueur de la liste = %d\n", liste.Length())


// test la fonction remove le cas l'elm est millieu de la liste 
	fmt.Println("test remove avec elm au milieu ")
	liste.Remove(12)
	fmt.Println("liste =", liste)

//teste cas elt est en tête de la liste 
	liste.Remove(2)
	fmt.Println( liste)
//teste cas elt est en queue de la liste 
	liste.Remove(15)
	fmt.Println( " la liste après la supression de dernier de la liste" ,liste)
	

// test de nouveau la fonction length 
	fmt.Printf("Longueur de la liste = %d\n", liste.Length())

// test la fonction fromSlice :
	liste3 := listedbchainee.FromSlice([]int{10, 20, 3, 1})
	fmt.Println(liste3)
	fmt.Println(liste3.ReverseString())



// test fonction append 	
	/*fmt.Println("test append ")
	liste.Append(100)
	liste.Append(200)
	fmt.Println("liste =" ,liste)
	fmt.Println(liste.ReverseString())
	*/

	liste4 := listedbchainee.FromSlice([]int{10})
	liste4.Remove(10)
	fmt.Println(liste4)
}

