package dict

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// URL récupération du dictionnaire en ligne
const url string = "https://dibi-dictionary.herokuapp.com/dictionary/getWords/all"

// Dictionnaire
var Dictionary []DibiWord

// Récupération du dictionnaire local
func GetDict() {
	fmt.Println("Récupération du dictionnaire sur dibi-dictionary.")

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erreur : %s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Erreur : %s", err)
	}

	// Parse en json
	err = json.Unmarshal([]byte(body), &Dictionary)
	if err != nil {
		fmt.Println("Error parsing JSON", err)
	}
}
