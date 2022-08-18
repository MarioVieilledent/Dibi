package dict

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// URL récupération du dictionnaire en ligne
const url string = "https://dibi-dictionary.herokuapp.com/dictionary/getWords/all"

// Chemin et nom du fichier JSON du dictionnaire local
var dictFilePath string = "./data/dict.json"

// Dictionnaire
var Dictionary []DibiWord

// Récupération du dictionnaire local
func GetDict() {
	// Si le fichier n'existe pas, on le créé avec les données récupérées en ligne
	if _, err := os.Stat(dictFilePath); errors.Is(err, os.ErrNotExist) {
		WriteDict()
	}

	// Lecture du fichier
	d, err := ioutil.ReadFile(dictFilePath)
	if err != nil {
		fmt.Println("Error reading local file", dictFilePath, ":", err)
	}

	// Parse en json
	err = json.Unmarshal(d, &Dictionary)
	if err != nil {
		fmt.Println("Error parsing JSON", err)
	}
}

// Récupération du dictionnaire en ligne et écriture dans le fichier local
func WriteDict() {
	fmt.Println("Récupération du dictionnaire sur dibi-dictionary.")

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erreur : %s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Erreur : %s", err)
	}

	sb := string(body)

	err = ioutil.WriteFile(dictFilePath, []byte(sb), 0777)

	if err != nil {
		fmt.Printf("Error writing file : %s", err)
	}
}

// Getter fichier JSON local
func GetDictFilePath() string {
	return dictFilePath
}

// Setter fichier JSON local
func SetDictFilePath(path string) {
	dictFilePath = path
}
