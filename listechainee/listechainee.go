package listechainee

import (
	"strconv"
)

// Cellule permet de representer un entier chainé
type ChainedInt struct {
	Valeur  int
	NextInt *ChainedInt
}

func Append(value int, head *ChainedInt) {
	var newCel ChainedInt
	newCel.Valeur = value
	newCel.NextInt = head
	head = &newCel
}

func (head *ChainedInt) String() string {
	var returnValue string = "{"
	for head != nil {
		returnValue += strconv.Itoa(head.Valeur)
		
		head = head.NextInt
		if head != nil{
			returnValue += ", "
		}

	}
	returnValue += "}\n"
	return returnValue
}
