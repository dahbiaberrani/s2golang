package pileslice

import(
	"fmt"
	"errors"
)

type Pile struct {
	corps [] int
	}

func CreerPile() Pile {
	return Pile{}
}

func (p Pile) String() string {
	res := "{"
	if !p.EstVide() {
		for i := range p.corps {
			res += fmt.Sprintf("%d", p.corps[i]) + ", "
		}
	}
	if res != "{" {
		return res[:len(res)-2]+"}"
	} else {
		return res+"}"
	}
}

func (p Pile) EstVide() bool{
	return len(p.corps) == 0
}

func (p *Pile) Empiler(elm int){
	p.corps = append(p.corps, elm)
}

func (p Pile) Sommet() (int, error){
	if p.EstVide() {
		return 0, errors.New("Impossible, la pile est vide")
	} else {
		return p.corps[len(p.corps) -1], nil
	}
}


func (p *Pile) Depiler() (int, error) {

	sommet, err := p.Sommet()
	if err == nil {
		p.corps = p.corps[:len(p.corps) -1 ]
	}	
	return sommet, err

}

func  FromSlice(slice [] int) Pile{
	pile := CreerPile()
	if len(slice) > 0 {
		for i := range slice {
			pile.Empiler(slice[i])
		}
	}
	return pile
}

func (p Pile) ToSlice() [] int{
	return p.corps
}
