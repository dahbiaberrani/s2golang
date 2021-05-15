package dblelinkedlist

import "fmt"

type cellule struct {
	valeur        int
	prec, suivant *cellule
}

func (cell cellule) String() string {
	return fmt.Sprintf("%d", cell.valeur)
}

// DbleLinkedList implémente une liste doublement chaînée
type DbleLinkedList struct {
	tete, queue *cellule
	nbElts      int
}

// String produit une représentation textuelle de la liste
func (list DbleLinkedList) String() string {
	res := ""
	for curr := list.tete; curr != nil; curr = curr.suivant {
		res += curr.String() + " "
	}
	if res != "" {
		return res[:len(res)-1] // Pour supprimer le dernier espace
	} else {
		return ""
	}
}

// Length renvoie le nombre d'éléments dans la liste
func (list DbleLinkedList) Length() int {
	return list.nbElts
}

// IsEmpty indique si la liste est vide
func (list DbleLinkedList) IsEmpty() bool {
	return list.nbElts == 0
}

// Append ajoute un élément à la fin de la liste
func (list *DbleLinkedList) Append(elt int) {
	nouv := &cellule{valeur: elt} // prec et suivant valent nil par défaut
	if list.IsEmpty() {           // Premier ajout...
		list.tete = nouv
		list.queue = list.tete
	} else { // On ajoute après list.queue...
		nouv.prec = list.queue // sans oublier le chainage précédent
		list.queue.suivant = nouv
		list.queue = nouv
	}
	list.nbElts++
}

// Cons ajoute un élément en tête de la liste
func (list *DbleLinkedList) Cons(elt int) {
	nouv := &cellule{valeur: elt, prec: nil, suivant: list.tete}
	list.nbElts++

	if list.IsEmpty() { // Premier ajout...
		list.queue = nouv
	} else { // Ajoute avant list.tete
		list.tete.prec = nouv
	}
	list.tete = nouv
}

// First renvoie l'adresse de la cellule contenant la première occurrence de elt, nil sinon
func (list *DbleLinkedList) First(elt int) *cellule {
	curr := list.tete
	for curr != nil && curr.valeur != elt {
		curr = curr.suivant
	}
	return curr
}

// Delete supprime la première occurrence de elt, sans effet sinon
func (list *DbleLinkedList) Delete(elt int) {
	toDelete := list.First(elt)
	if toDelete != nil {
		if toDelete == list.tete { // suppression en tête
			list.tete = toDelete.suivant
			if list.tete != nil {      // Cas où la liste a plus d'un élément
				list.tete.prec = nil
			} else {                   // Cas où la liste n'avait qu'un seul élement
				list.queue = nil
			}
		} else {
			toDelete.prec.suivant = toDelete.suivant
			if toDelete.suivant != nil {                 // suppression "au milieu"
				toDelete.suivant.prec = toDelete.prec
			} else {                                     // suppression "en queue"
				list.queue = toDelete.prec
			}
		}
		list.nbElts--
	}
}

// FromList crée une DbleLinkedList à partir d'une tranche (éventuellement vide)
func FromList(tab []int) DbleLinkedList {
	res := DbleLinkedList{} // tete et queue valent nil par défaut et nbElts vaut 0
	for _, elt := range tab {
		res.Append(elt) // Ajoute à la fin pour préserver l'ordre
	}
	return res
}

// ToList crée un slice à partir de la liste
func (list DbleLinkedList) ToList() []int {
	res := make([]int, list.nbElts) // Crée une tranche avec le bon nombre d'éléments
	i := 0
	for curr := list.tete; curr != nil; curr = curr.suivant {
		res[i] = curr.valeur
		i++
	}
	return res
}
