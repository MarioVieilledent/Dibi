package httpService

import "os"

var Url string = "0.0.0.0"

// Setup du port dans env pour Heroku
var port string = getPort()

func getPort() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return
}
