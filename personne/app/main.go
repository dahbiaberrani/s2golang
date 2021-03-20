package main

import (
	"fmt"
	"personne"
	"strings"
)

func main() {
	var personnes []personne.Personne // une tranche de personnes
	// saisie des personnes
	encore := true
	for encore {
		rep := ""
		personnes = append(personnes, personne.SaisiePersonne())
		fmt.Print("encore (o/n)?")
		fmt.Scanln(&rep)
		encore = strings.HasPrefix(strings.ToUpper(rep), "O")

	}
	// Affichage des toues les personnes stockées dans la tranche
	for _, pers := range personnes {
		fmt.Println(pers)
	}

}
