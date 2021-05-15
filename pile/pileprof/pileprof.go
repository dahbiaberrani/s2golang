package  pileprof

import "fmt"

type Cellule struct {
	valeur  int
	next *Cellule
}



// Pile est implantée comme une liste chainée
type Pile struct {
	sommet *Cellule
}

// Méthode de conversion d'une cellule en chaîne : redéfinition de la méthode String()
func (c Cellule) String() string {
	return fmt.Sprintf("%d", c.valeur)
}

func (p Pile) String() string {
	res := "{"
	curr := p.sommet
	for curr != nil {
		res += curr.String() + "-> "
		curr = curr.next
	}
	if res != "{" {
		return res[:len(res)-3]+"}"
	} else {
		return res+"}"
	}
}

func CreerPile() Pile {
	return Pile{nil}
}

// FromSlice crée une Pile à partir d'une tranche. Le sommet est le premier elt
func FromSlice (elts ... int) Pile {
	res := Pile{}
	for _, e := range elts {
		res.Push(e)
	}
	return res
}

// IsEmpty teste si la Pile est vide
func (p Pile) IsEmpty() bool {
	return p.sommet == nil
}

// Push ajoute une élément au sommet de la Pile
func (p *Pile) Push(e int){
	p.sommet = &Cellule{e, p.sommet}
}

// Pop supprime et renvoie l'élement au sommet
func (p *Pile) Pop() (int, error) {
	if p.IsEmpty(){
		return 0, fmt.Errorf("empty stack")
	}
	res := p.sommet.valeur
	p.sommet = p.sommet.next
	return res, nil
}

// ToSlice produit une tranche à partir d'une pile
func (p Pile) ToSlice() [] int{
	res := []int{}
	for curr :=p.sommet; curr !=nil; curr = curr.next{
		res = append(res, curr.valeur)
	}
	return res
}

// Top renvoie l'élement au sommet de la Pile
func (p Pile) Top() (int, error) {
	if p.IsEmpty() {
		return 0, fmt.Errorf("empty stack")
	}
	return p.sommet.valeur, nil
}