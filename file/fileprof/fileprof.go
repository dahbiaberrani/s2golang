package fileprof

import "fmt"

type cellule struct {
	val		int
	next	*cellule
}

type File struct {
	tete	*cellule
	queue   *cellule
	nbElts	int
}

func CreerFile() File {
	return File{nil,nil,0}
}

// Méthode de conversion d'une cellule en chaîne : redéfinition de la méthode String()
func (c cellule) String() string {
	return fmt.Sprintf("%d", c.val)
}

func (f File) String() string {
	res := "{"
	curr := f.tete
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

//FromSlice crée une file à partir d'une tranche
func FromSlice(tab ...int) File {
	res := File{}
	for i :=0; i < len(tab); i++ {
		res.Ajouter(tab[i])
	}
	return res
}

//ToSlice renvoie une tranche représentant la file
func (f File) ToSlice() []int {
	res := []int{}
	for curr := f.tete; curr !=nil; curr = curr.next{
		res = append(res, curr.val)
	}
	return res
}

//EstVide teste si la file est vide
func (f File) EstVide() bool {
	return f.nbElts == 0
}

func (f File) NbElts() int{
	return f.nbElts
}

//Ajouter ajoute elt a la fin de la file
func (f *File) Ajouter(elt int) {
	nouv := &cellule{elt,nil}
	if !f.EstVide() {
		f.queue.next = nouv
		
	} else {
		f.tete = nouv
	}
	
	f.queue = nouv
	f.nbElts++
}

func (f *File) Oter() (int, error) {

	tete, err := f.Tete()
	if err == nil {
		f.tete = f.tete.next
		f.nbElts--
	}	
	return tete, err
}

//Tete renvoie l'element en tête de la file
func (f File) Tete() (int, error) {
	if f.nbElts == 0 {
		return 0, fmt.Errorf("Impossible, File Vide")
	} else {
		return f.tete.val, nil
	}
}