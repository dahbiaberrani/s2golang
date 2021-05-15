
package pilechainee
import(
 "listechainee"
 "errors"
)

type Pile struct {
	corps listechainee.Liste
	}

func CreerPile() Pile {
	return Pile{}
}

func (p Pile) String() string {
	return "{"+p.corps.String()+"}"
}

func (p Pile) EstVide() bool{
	return p.corps.Empty()
}

func (p *Pile) Empiler(elm int){
	p.corps.Cons(elm)
}

func (p Pile) Sommet() (int, error){
	if p.EstVide() {
		return 0, errors.New("Impossible, la pile est vide")
	} else {
		tab := p.ToSlice()
		return tab[0], nil
	}
}


func (p *Pile) Depiler() (int, error) {

	sommet, err := p.Sommet()
	if err == nil {
		p.corps = listechainee.FromSlice(p.ToSlice()[1:])
	}	
	return sommet, err
}

func  FromSlice(slice [] int) Pile{
	pile := CreerPile()
	for i := range slice {
		pile.Empiler(slice[i])
	}
	return pile
}

func (p Pile) ToSlice() [] int{
	return p.corps.ToSlice()
}
