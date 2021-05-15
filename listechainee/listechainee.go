package listechainee
import (
	"fmt"
)
type cellule struct {
	valeur  int
	suivant *cellule
}

// Méthode de conversion d'une cellule en chaîne : redéfinition de la méthode String()
func (c cellule) String() string {
	return fmt.Sprintf("%d", c.valeur)
}

// Liste est un type liste simplement chaînée
type Liste struct {
	tete, queue   *cellule
	nbElts int
}

// String fournit une représentation textuelle de la liste
func (ls Liste) String() string {
	res := ""
	curr := ls.tete
	for curr != nil {
		res += curr.String() + "-> "
		curr = curr.suivant
	}
	if res != "" {
		return res[:len(res)-3]
	} else {
		return res
	}
}

// Cons ajoute un elément en tête de la liste
func (ls *Liste) Cons(elt int) {
	// Créer une nouvelle cellule avec elt dans valeur et ls.tete dans suivant
	// et faire pointer ls.tete dessus
	ls.tete = &cellule{elt, ls.tete}
	ls.nbElts++
	if ls.queue == nil {
	   ls.queue = ls.tete
	}
}

// Length renvoie le nombre d'éléments dans la liste
func (ls Liste) Length() int {
	return ls.nbElts
}

// Empty teste si la liste est vide
func (ls Liste) Empty() bool {
	return ls.tete == nil
}

// Append ajoute elt à la fin de la liste
func (ls *Liste) Append(elt int) {
	nouv := &cellule{elt, nil}
	ls.nbElts++

	// Cas particulier : ls était vide
	if ls.tete == nil {
		ls.tete = nouv
		ls.queue = nouv
		return
	}

	ls.queue.suivant = nouv
	ls.queue = nouv
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
	
	// Cas particulier 3 : l'elt est le dernier de la liste
	// à faire pour la prochaine fois...
	
	// Cas général
	curr := ls.tete
	for curr.suivant != nil && curr.suivant.valeur != elt {
		curr = curr.suivant
	}
	if curr.suivant != nil {
		curr.suivant = curr.suivant.suivant
		ls.nbElts--
	}
}