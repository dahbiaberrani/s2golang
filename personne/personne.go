package personne
import (
	"fmt"
	"bufio"
	"os"
)


// Adresse permet de représenter un adresse postal 	
type Adresse struct{
	Rue, Ville, Cp string
}


// Personne permet de représenter une personne
type Personne struct {
	Prenom, Nom, DateNaissance string 
	Adresse						Adresse
}

// Redéfinition de la représentation textuelle d'une Adresse


func (a Adresse) String() string {
	return fmt.Sprintf("%s, %s %s", a.Rue, a.Cp, a.Ville)
	}
// Redéfinition de la représentation textuelle d'une personne

func (p Personne) String() string {
	return fmt.Sprintf("%s %s, né le %s. Adresse : %s.", p.Prenom, p.Nom, p.DateNaissance, p.Adresse)

}

func SaisiePersonne() Personne {
	var (
		res  	Personne
		scanner = bufio.NewScanner(os.Stdin)
	 )
	fmt.Print("Prénom : ")
	scanner.Scan()
	res.Prenom = scanner.Text()
	fmt.Print("Nom : ")
	scanner.Scan()
	res.Nom = scanner.Text()
	fmt.Print("Date de naissance (JJ/MM/AAAA) : ")
	scanner.Scan()
	res.DateNaissance = scanner.Text()
	res.Adresse = saisieAdresse()
	return res
}
func saisieAdresse() Adresse {
	var (
		res  	Adresse 
		scanner = bufio.NewScanner(os.Stdin)
	 )
	fmt.Print("Rue : ")
	scanner.Scan()
	res.Rue = scanner.Text()
	fmt.Print("Ville : ")
	scanner.Scan()
	res.Ville = scanner.Text()
	fmt.Print("Code postal : ")
	scanner.Scan()
	res.Cp = scanner.Text()
	return res
}