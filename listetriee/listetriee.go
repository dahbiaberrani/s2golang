
package listetriee // import "listetriee"
import (
	//"strconv"
)

//TYPES
type cellule struct{
	valeur  int
	suivant  *cellule

}
type ListeTriee struct {
	head *cellule
	tail *cellule
	
}
   


   



/*
func (list *ListeTriee) Delete(val int)
    //Pas demandé mais, à y être... Delete supprime la première occurrence de lavaleur val dans la liste

func (list ListeTriee) First(elt int) *cell
    //First renvoie la première cellule contenant la valeur elt dans la liste (nil si la liste est vide)



func (list *ListeTriee) Replace(old, nouv int)
    //Replace remplace la première occurrence de la valeur old par la valeur nouv dans la liste
*/

/*
func (list ListeTriee) ToList() []int
    //ToList renvoie une tranche contenant les éléments de la liste
	*/