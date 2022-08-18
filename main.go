package main

import (
	"Dibi/src/dict"
	"Dibi/src/httpService"
)

func main() {
	// Récupération du dictionnaire
	go dict.GetDict()

	// Lancement de l'API
	go httpService.LaunchAPI()

	// Service statique de la page web
	httpService.ServeWebApp()
}
