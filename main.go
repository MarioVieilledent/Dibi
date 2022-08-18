package main

import (
	"Dibi/src/dict"
	"Dibi/src/httpService"
	"fmt"
)

func main() {

	a := "jiksaifo"
	b := "jiksai"

	for i := 0; i < len(a)+1; i++ {
		for j := i + 2; j < len(a)+1; j++ {
			for k := 0; k < len(b)+1; k++ {
				for l := k + 2; l < len(b)+1; l++ {
					if a[i:j] == b[k:l] {
						fmt.Println(a[i:j], "et", b[k:l])
					}
				}
			}
		}
	}

	// Récupération du dictionnaire
	go dict.GetDict()

	// Lancement de l'API
	go httpService.LaunchAPI()

	// Service statique de la page web
	httpService.ServeWebApp()
}
