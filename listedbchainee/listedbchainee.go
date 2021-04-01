package main

import (
	"fmt"
)

type cellule struct {
	valeur  int
	suivant *cellule
	precedant *cellule
}

// Méthode de conversion d'une cellule en chaîne : redéfinition de la méthode String()
func (c cellule) String() string {
	return fmt.Sprintf("%d", c.valeur)
}

// Liste est un type liste simplement chaînée
type Liste struct {
	tete,queue *cellule
	nbElts int 
}

// String fournit une représentation textuelle de la liste
func (ls Liste) String() string {
	res := "{"
	curr := ls.tete
	for curr != nil {
	//1.traitement de la cellule courante :
		res +=  curr.String()+"->"
	//2.passer à la cellule suivante
		curr = curr.suivant		
	}
	res += "nil}"
	return res
}
// Cons ajoute un elément en tête de la liste
func (ls *Liste) Cons(elt int) {
	nouv := &cellule{elt , ls.tete , ls.queue}
	ls.tete = nouv	
	nouv = nil	
	ls.nbElts ++
}

/*
// Length renvoie le nombre d'éléments dans la liste
func (ls Liste) Length() int {

	return  nbElts
}

// Empty teste si la liste est vide
func (ls Liste) Empty() bool {
	return ls.tete == nil
		
}


// Append ajoute elt à la fin de la liste
func (ls *Liste) Append(elt int) {
	nouv := &cellule{elt,ls.queue}
	ls.queue = nouv	
	ls.nbElts++
		
}
// Remove supprime un élément dans la liste
func (ls *Liste) Remove(elt int) {
		// Cas particulier 1 : la liste est vide
	if ls.tete == nil {
		return
	}
		// Cas particulier 2 : l'elt  est en tete
	if ls.tete.valeur == elt {
		ls.tete = ls.tete.suivant
		ls.tete.suivant = nil
		return
		}
	// cas générale 
	curr := ls.tete
	prec := nil
	for curr != nil && curr.valeur != elt {
		prec = curr
		curr = curr.suivant 
		
	}
	if (curr != nil && cur.valeur == elt){
		prec = curr.suivant 
		cur.suivant = prec
	}
}


*/



func main() {
	liste := Liste{}
	fmt.Println(liste)
	liste.Cons(12)
	liste.Cons(42)
	liste.Cons(55)
	fmt.Println(liste)
	/*fmt.Printf("Longueur de la liste = %d\n", liste.Length())

	liste.Append(100)
	liste.Append(200)
	fmt.Println(liste)

	liste2 := Liste{}
	liste2.Remove(100)
	fmt.Println("liste2 =", liste2)



	liste3.Remove(10)
	fmt.Println(liste3)
	liste3.Remove(3)
	fmt.Println(liste3)
	liste3.Remove(42)
	fmt.Println(liste3)
*/
}