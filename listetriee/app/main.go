package main

import (
	"fmt"
	//"listetriee"
	"strconv"
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

func (list ListeTriee) String() string{
    //String renvoie une représentation textuelle de la liste

	stringValue :="{"
	courant := list.head
	for courant != nil{
		//1. traitement de la cellule courante :
		stringValue += strconv.Itoa(courant.valeur)+" ->"
		//2. passer à la cellule suivante 
		courant = courant.suivant
	}	
	stringValue +="nil}"	
	return stringValue
}

//Add ajoute un élément *à sa bonne place* dans la liste
func (list *ListeTriee) Add(elt int){
	// cas de la liste vide
	nouvcellule := cellule{valeur:elt}
	nouvcellule.valeur = elt
	if  list.head == nil {
		list.head = &nouvcellule
		list.tail = &nouvcellule
		nouvcellule.suivant = nil
	}else{
		nouvcellule := cellule{valeur:elt}
		//cas d'insertion au debut 
		if elt <= list.head.valeur{	
			nouvcellule.suivant = list.head
			list.head = &nouvcellule
		}else{
			// cas d'insertion à la fin 

			if elt >= list.tail.valeur{
				list.tail.suivant = &nouvcellule
				list.tail = &nouvcellule
				nouvcellule.suivant= nil
			}else{
				//cas d'insertion au milieu

				//recherche de la place
				precedcourant := list.head
				courant := list.head
				for elt > courant.valeur{
					precedcourant = courant
					courant = courant.suivant
				}
				nouvcellule.suivant = courant
				precedcourant.suivant= &nouvcellule
			}
		}
	}
}

  //ListeTriee implémente une liste d'entiers en permanence triée par ordre croissant
func(ListeTriee *ListeTriee)FromList(goList []int) {
	
	//1. Ajout un par un les element du tableau en parcourant
		for i:=0; i < len(goList);i++{
			ListeTriee.Add(goList[i])
		}
}	
func (list ListeTriee) Length() int{
	var compteur int = 0
	courant := list.head
	for courant != nil {
		compteur +=1
		courant = courant.suivant
	}
	return compteur
}
func (list ListeTriee) ToList() []int{
    //ToList renvoie une tranche contenant les éléments de la liste
	res := []int{}
	courant := list.head
	for courant != nil{
		res= append(res,courant.valeur)
		courant=courant.suivant
	}
	return res
}
func (list *ListeTriee) Delete(elt int)bool{
		// Cas particulier 1 : la liste est vide
	if list.head == nil {
		return false
	}
	// cas elt est au debut de la liste 

	if elt == list.head.valeur {		
		list.head=list.head.suivant
		return true
	}
	//cas elt est à la fin de liste 
	
	//cas génerale 
	courant := list.head
	for courant.suivant != nil && elt != courant.suivant.valeur{
		courant = courant.suivant
		
	}
	if courant.suivant != nil{
		courant.suivant = courant.suivant.suivant
		return true 

	}
	return false

	
}

func (list *ListeTriee)Replace(old,nouv int){
	
	for list.Delete(old) {

		list.Add(nouv)
	}
}


func main() {
	/*
	var val int
	fmt.Print("Entrez une valeur : ")
	fmt.Scan(&val)                      // On entre 42

	if maListeTriee.First(val) != nil {
		fmt.Printf("%d est dans la liste\n", val)   // 42 est dans la liste
	} else {
		fmt.Printf("%d n'est pas dans la liste\n", val)
	}

	
	fmt.Println(maListeTriee)              // 1 12 15 42 54 78
	maListeTriee.Replace(1, 13)
	fmt.Println(maListeTriee)              // 12 13 15 42 54 78
	maListeTriee.Replace(42, 1)
	fmt.Println(maListeTriee)              // 1 12 13 15 54 78
	*/

	ma_liste :=  ListeTriee{}
	ma_liste.head = nil
	ma_liste.tail = nil


	ma_liste.Add(45)
	ma_liste.Add(40)
	ma_liste.Add(50)
	ma_liste.Add(55)
	ma_liste.Add(65)
	ma_liste.Add(20)
	ma_liste.Add(41)
	ma_liste.Add(62)
	ma_liste.Add(23)
	ma_liste.Add(5)

	fmt.Println(ma_liste)
	maListeTriee := ListeTriee{head:nil,tail:nil}
	maListeTriee.FromList([]int{12, 1, 78, 42, 15})
	fmt.Println(maListeTriee)
	fmt.Println(maListeTriee.Length())     // 5
	maListe := maListeTriee.ToList()
	fmt.Println(maListe)                   // [1 12 15 42 78]



	fmt.Println("=============================")
	maListe1 := ListeTriee{head:nil,tail:nil}
	fmt.Println(maListe1)
	maListe1.Delete(1)
	fmt.Println(maListe1)

	maListe1.Add(12)
	fmt.Println(maListe1)
	maListe1.Delete(1)
	fmt.Println(maListe1)
	maListe1.Delete(12)
	fmt.Println(maListe1)


	maListe1.FromList([]int{12, 1, 78, 42, 15})
	fmt.Println(maListe1)
	maListe1.Delete(1)
	fmt.Println(maListe1)
	maListe1.Delete(15)
	fmt.Println(maListe1)
	maListe1.Delete(78)
	fmt.Println(maListe1)
	maListe1.Add(12)
	maListe1.Add(12)
	maListe1.Add(12)
	fmt.Println(maListe1)

	maListe1.Replace(12, 13)
	fmt.Println(maListe1)


	
}