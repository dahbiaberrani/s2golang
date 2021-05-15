package filechainee
import (
	"fmt"
	"listechainee"

)


type File struct {
	corps listechainee.Liste
	}

func CreerFile() File {
	return File{}
}

func (p File) NbElts() int{
	return p.corps.Length()
}

func (p File) String() string {
	return "{"+p.corps.String()+"}"
}

func (p File) EstVide() bool{
	return p.corps.Empty()
}

func (p *File) Ajouter(elm int){
	p.corps.Append(elm)
}

func (p File) Tete() (int, error){
	if p.EstVide() {
		return 0, fmt.Errorf("Impossible, la pile est vide")
	} else {
		return p.ToSlice()[0], nil
	}
}


func (p *File) Oter() (int, error) {

	sommet, err := p.Tete()
	if err == nil {
		p.corps = listechainee.FromSlice(p.ToSlice()[1:])
	}	
	return sommet, err
}

func  FromSlice(slice [] int) File{
	file := CreerFile()
	for i := range slice {
		file.Ajouter(slice[i])
	}
	return file
}

func (p File) ToSlice() [] int{
	return p.corps.ToSlice()
}
