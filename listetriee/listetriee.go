
package listetriee // import "listetriee"
import (
	"strconv"
)

//TYPES
type cellule struct{
	valeur  int
	suivant  *cellule

}
type ListeTriee struct {
	head *cellule
	tail *cellule
	
}

func (list ListeTriee) String() string{
    //String renvoie une représentation textuelle de la liste

	stringValue :="{"
	courant := list.head
	for courant != nil{
		//1. traitement de la cellule courante :
		stringValue += strconv.Itoa(courant.valeur)+" ->"
		//2. passer à la cellule suivante 
		courant = courant.suivant
	}	
	stringValue +="nil}"	
	return stringValue
}

//Add ajoute un élément *à sa bonne place* dans la liste
func (list *ListeTriee) Add(elt int){
	// cas de la liste vide
	nouvcellule := cellule{valeur:elt}
	nouvcellule.valeur = elt
	if  list.head == nil {
		list.head = &nouvcellule
		list.tail = &nouvcellule
		nouvcellule.suivant = nil
	}else{
		nouvcellule := cellule{valeur:elt}
		//cas d'insertion au debut 
		if elt <= list.head.valeur{	
			nouvcellule.suivant = list.head
			list.head = &nouvcellule
		}else{
			// cas d'insertion à la fin 

			if elt >= list.tail.valeur{
				list.tail.suivant = &nouvcellule
				list.tail = &nouvcellule
				nouvcellule.suivant= nil
			}else{
				//cas d'insertion au milieu

				//recherche de la place
				precedcourant := list.head
				courant := list.head
				for elt > courant.valeur{
					precedcourant = courant
					courant = courant.suivant
				}
				nouvcellule.suivant = courant
				precedcourant.suivant= &nouvcellule
			}
		}
	}
}

  //ListeTriee implémente une liste d'entiers en permanence triée par ordre croissant
func(ListeTriee *ListeTriee)FromList(goList []int) {
	
	//1. Ajout un par un les element du tableau en parcourant
		for i:=0; i < len(goList);i++{
			ListeTriee.Add(goList[i])
		}
}	
func (list ListeTriee) Length() int{
	var compteur int = 0
	courant := list.head
	for courant != nil {
		compteur +=1
		courant = courant.suivant
	}
	return compteur
}
func (list ListeTriee) ToList() []int{
    //ToList renvoie une tranche contenant les éléments de la liste
	res := []int{}
	courant := list.head
	for courant != nil{
		res= append(res,courant.valeur)
		courant=courant.suivant
	}
	return res
}
func (list *ListeTriee) Delete(elt int)bool{
		// Cas particulier 1 : la liste est vide
	if list.head == nil {
		return false
	}
	// cas elt est au debut de la liste 

	if elt == list.head.valeur {		
		list.head=list.head.suivant
		return true
	}
	//cas elt est à la fin de liste 
	
	//cas génerale 
	courant := list.head
	for courant.suivant != nil && elt != courant.suivant.valeur{
		courant = courant.suivant
		
	}
	if courant.suivant != nil{
		courant.suivant = courant.suivant.suivant
		return true 

	}
	return false

	
}

func (list *ListeTriee)Replace(old,nouv int){
	
	for list.Delete(old) {

		list.Add(nouv)
	}
}