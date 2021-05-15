package listetriee_prof

import "fmt"

type cell struct {
	val  int
	suiv *cell
}

func (c cell) String() string {
	return fmt.Sprintf("%d", c.val)
}

// ListeTriee implémente une liste d'entiers en permanence triée par ordre croissant
type ListeTriee struct {
	tete   *cell
	nbElts int
}

// String renvoie une représentation textuelle de la liste
func (list ListeTriee) String() string {
	res := ""
	for cur := list.tete; cur != nil; cur = cur.suiv {
		res += cur.String() + " "
	}
	if res != "" {
		return res[:len(res)-1]
	} else {
		return ""
	}
}

// Length renvoie le nombre d'éléments de la liste
func (list ListeTriee) Length() int {
	return list.nbElts
}

// Add ajoute un élément *à sa bonne place* dans la liste
func (list *ListeTriee) Add(elt int) {
	nouv := &cell{elt, list.tete}

	if list.tete == nil || list.tete.val > elt { // Premier ajout...
		list.tete = nouv
	} else {
		cur := list.tete
		for cur.suiv != nil && cur.suiv.val < elt {  // On regarde une cellule "vers l'avant"
			cur = cur.suiv
		}
		nouv.suiv = cur.suiv
		cur.suiv = nouv
	}
	list.nbElts++
}

// FromList crée une ListeTriee à partir d'une tranche
func FromList(goList []int) ListeTriee {
	res := ListeTriee{}
	for _, elt := range goList {
		res.Add(elt)
	}
	return res
}

// ToList renvoie uns tranche contenant les éléments de la liste
func (list ListeTriee) ToList() []int {
	res := make([]int, list.nbElts)
	i := 0
	for cur := list.tete; cur != nil; cur = cur.suiv {
		res[i] = cur.val
		i++
	}

	/*res := []int{}
	for cur := list.tete; cur != nil; cur = cur.suiv {
		res = append(res, cur.val)
	}*/

	return res
}

// First renvoie la première cellule contenant la valeur elt dans la liste
// (nil si la liste est vide ou si elt n'est pas dans la liste)
func (list ListeTriee) First(elt int) *cell {
	cur := list.tete
	for cur != nil && cur.val < elt {
		cur = cur.suiv
	}
	if cur != nil && cur.val == elt {
		return cur
	} else {
		return nil
	}
}

// Replace remplace la première occurrence de la valeur old par la valeur nouv dans la liste
func (list *ListeTriee) Replace(old, nouv int) {
	pos := list.First(old)
	if  pos != nil {
		pos.val = nouv
		tmp := list.ToList()
		*list = FromList(tmp)   // FromList (donc Add) remet tout en ordre
	}
}

// Pas demandé mais, à y être...
// Delete supprime la première occurrence de la valeur val dans la liste
func (list *ListeTriee) Delete(val int) {
	if list.tete == nil || val < list.tete.val { // val n'est donc pas dans la liste
		return               // on ne fait rien...
	}
	cur := list.tete
	if cur.val == val {      // Suppression en tête
		list.tete = list.tete.suiv
		list.nbElts--
		return
	}
	// Cas général : parcours de la liste "une cellule vers l'avant"...
	for cur.suiv != nil && cur.suiv.val < val {
		cur = cur.suiv
	}
	if cur.suiv != nil && cur.suiv.val == val {     // Trouvé
		cur.suiv = cur.suiv.suiv                    // on supprime
		list.nbElts--
	}
}
