package dict

type DibiWord struct {
	Id           string `json:"_id"`
	Dibi         string `json:"dibi"`
	French       string `json:"french"`
	English      string `json:"english"`
	PartOfSpeech string `json:"partOfSpeech"`
	Author       string `json:"author"`
	Date         string `json:"date"`
	Description  string `json:"description"`
}
