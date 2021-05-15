package fileslice 

import(
	"fmt"
)

type File struct {
	corps [] int
	}

func CreerFile() File {
	return File{}
}

func (p File) String() string {
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

func (p File) NbElts() int{
	return len(p.corps)
}

func (p File) EstVide() bool{
	return len(p.corps) == 0
}

func (p *File) Ajouter(elm int){
	p.corps = append(p.corps, elm)
}

func (p File) Tete() (int, error){
	if p.EstVide() {
		return 0, fmt.Errorf("Impossible, la pile est vide")
	} else {
		return p.corps[0], nil
	}
}


func (p *File) Oter() (int, error) {

	sommet, err := p.Tete()
	if err == nil {
		p.corps = p.corps[1:]
	}	
	return sommet, err

}

func  FromSlice(slice [] int) File{
	file := CreerFile()
	if len(slice) > 0 {
		for i := range slice {
			file.Ajouter(slice[i])
		}
	}
	return file
}

func (p File) ToSlice() [] int{
	return p.corps
}
