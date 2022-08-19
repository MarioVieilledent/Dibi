package httpService

import "os"

var Url string = "127.0.0.1"

var APIPort string = "3001"

// Setup du port dans env pour Heroku
var port string = getPort()

func getPort() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return
}
