
// BERRANI Dahbia
// Examen go semestre 2
package listetriee

import (
	"strconv"

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

func (list Ensemble) String() string{
    //String renvoie une représentation textuelle de la liste

	stringValue :="{"
	courant := list.tete
	for courant != nil{
		//1. traitement de la Cellule courante :
		stringValue += strconv.Itoa(courant.valeur)+" ->"
		//2. passer à la Cellule suivante 
		courant = courant.suivant
	}	
	stringValue +="nil}"	
	return stringValue
}

//EstVide teste si la file est vide
func (e Ensemble) EstVide() bool {
	return f.cardinal == 0
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
func (e Ensemble)  Ajouter(elm int) {
	if !Appartient(elm) {

		
	}
}

