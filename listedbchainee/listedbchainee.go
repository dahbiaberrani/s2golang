package listedbchainee

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


// Length renvoie le nombre d'éléments dans la liste
func (ls Liste) Length() int {

	return  ls.nbElts
}

// Empty teste si la liste est vide
func (ls Liste) Empty() bool {
	return ls.tete == nil
		
}


// Append ajoute elt à la fin de la liste
func (ls *Liste) Append(elt int) {
	nouv := &cellule{elt, ls.queue ,ls.tete}

	// Cas particulier : ls était vide
	if ls.tete == nil {
		ls.tete = nouv
		ls.nbElts ++
		return
	}

	// Cas général : on se place sur la dernière cellule et on lui ajoute un lien vers nouv
	curr := ls.tete
	for curr.suivant != nil {
		curr = curr.suivant
	}
	curr.suivant = nouv
	ls.nbElts ++
		
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
		ls.nbElts--
		return
		}
	// cas générale 
	curr := ls.tete
	prec := curr
	for curr != nil && curr.valeur != elt {
		prec = curr
		curr = curr.suivant 
		
	}
	if (curr != nil && curr.valeur == elt){
		curr = curr.suivant
		prec.suivant = curr
		curr = prec.suivant
		ls.nbElts--
	}

}
// ToSlice crée une tranche à partir d'une liste
func (ls Liste) ToSlice() []int {
	res := []int{}
	curr := ls.tete
	for curr != nil {
		res = append(res, curr.valeur)
		curr = curr.suivant
	}
	return res
}

// FromSlice crée une liste à partir d'une tranche
func FromSlice(tab []int) Liste {
	res := Liste{}
	for i := len(tab) - 1; i >= 0; i-- {
		res.Cons(tab[i])
	}
	return res
}




