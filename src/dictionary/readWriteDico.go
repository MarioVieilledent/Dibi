package dictionary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// URL récupération du dictionnaire en ligne
const url string = "https://dibi-dictionary.herokuapp.com/dictionary/getWords/all"

// Chemin et nom du fichier JSON du dictionnaire local
var dicoFilePath string = "dico.json"

// Getter fichier JSON local
func GetDicoFilePath() string {
	return dicoFilePath
}

// Setter fichier JSON local
func SetDicoFilePath(path string) {
	dicoFilePath = path
}

// Récupération du dictionnaire local
func GetDico() (dic []DibiWord) {
	d, err := ioutil.ReadFile(dicoFilePath)

	if err != nil {
		fmt.Println("Error reading local file", dicoFilePath, ":", err)
	}

	err = json.Unmarshal(d, &dic)
	if err != nil {
		fmt.Println("Error parsing JSON", err)
	}

	return
}

// Récupération du dictionnaire en ligne et écriture dans le fichier local
func WriteDico() {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erreur : %s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Erreur : %s", err)
	}

	sb := string(body)

	err = ioutil.WriteFile(dicoFilePath, []byte(sb), 0777)

	if err != nil {
		fmt.Printf("Error writing file : %s", err)
	}
}
