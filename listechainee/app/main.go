package main

import (
	"fmt"
	"listechainee"

) 


func main(){
	var boite0 listechainee.ChainedInt 
	boite0.Valeur = 5
	boite0.NextInt = nil 
	var head *listechainee.ChainedInt = &boite0
	fmt.Print(head)
	listechainee.Append(12,head)
	fmt.Print(head)
	listechainee.Append(34,head)
	fmt.Print(head)




	
}
