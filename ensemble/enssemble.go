
// BERRANI Dahbia
// Examen go semestre 2
package ensemble

import (
	"fmt"

)

//TYPES
type Cellule struct{
	valeur  int
	suivant  *Cellule
}

type Ensemble struct {
	tete   *Cellule
	cardinal int
}

//String renvoie uen chane de caractères représentant l'ensemble e
func (e Ensemble) String() string{
    //String renvoie une représentation textuelle de la liste

	stringValue :="{"
	courant := e.tete
	for courant != nil{
		//1. traitement de la Cellule courante :
		stringValue += fmt.Sprintf("%d ", courant.valeur)
		//2. passer à la Cellule suivante 
		courant = courant.suivant
	}		
	return stringValue[:len(stringValue)-1]+"}"
}

//EstVide teste si la file est vide
func (e Ensemble) EstVide() bool {
	return e.cardinal == 0
}

//Card renvoi le cardinal de l'enssemble
func (e Ensemble) Card() int{
	return e.cardinal
}


// Appartient: renvoie vrais si l'element elm appartient a l'enssemble e sinon faux
func (e Ensemble) Appartient(elm int) bool {

	resultat := false
	curr := e.tete
	for curr != nil && !resultat {
		if curr.valeur == elm {
			resultat = true
		}
		curr = curr.suivant
	}

	return resultat
}

// Ajouter rajoute l'element elm a l'essemble e s'il n'existe pas déjà
func (e *Ensemble)  Ajouter(elm int) {
	if !e.Appartient(elm) {
		// Créer une nouvelle cellule avec elt dans valeur et e.tete dans suivant
		// et faire pointer e.tete dessus
		e.tete = &Cellule{elm, e.tete}
		e.cardinal++
	}
}

// FromSlice crée un ensemble à partir d'une tranche
func FromSlice(tab []int) Ensemble {
	res := Ensemble{nil,0}
	for i := 0; i < len(tab) ; i++ {
		res.Ajouter(tab[i])
	}
	return res
}



// Oter supprime un élément dans la liste
func (e *Ensemble) Oter(elt int) {
	// Cas particulier 1 : la liste est vide
	if e.tete == nil {
		return
	}

	// Cas particulier 2 : l'elt  est en tete
	if e.tete.valeur == elt {
		e.tete = e.tete.suivant
		e.cardinal--
		return
	}
	
	// Cas général
	curr := e.tete
	for curr.suivant != nil && curr.suivant.valeur != elt {
		curr = curr.suivant
	}
	if curr.suivant != nil {
		curr.suivant = curr.suivant.suivant
		e.cardinal--
		return
	}
}

// Inter renvoie l'intersection de deux ensembles
func (e Ensemble) Inter(ens Ensemble) Ensemble {
	res := Ensemble{nil,0}
	curr := e.tete
	for curr != nil {
		if ens.Appartient(curr.valeur)  {
			res.Ajouter(curr.valeur)
		}
		curr = curr.suivant
	}
	return res
}