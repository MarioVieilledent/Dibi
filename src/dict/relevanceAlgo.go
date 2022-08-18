package dict

import (
	"math"
	"sort"
	"strings"
)

const defaultListSize int = 50

func GetMatchingList(query string, listSize int) []DibiWord {

	sort.Slice(Dictionary, func(a, b int) bool {
		return Dictionary[a].calculateScore(query) > Dictionary[b].calculateScore(query)
	})

	if listSize >= len(Dictionary) {
		return Dictionary
	} else if listSize < 0 {
		return Dictionary[0:defaultListSize]
	} else {
		return Dictionary[0:listSize]
	}
}

func (dw DibiWord) calculateScore(query string) float64 {
	if dw.Dibi == query {
		return 1.0
	} else if dw.Dibi == "" || query == "" {
		return 0.0
	} else {
		score := 0.0
		score += stringsScore(dw.Dibi, query)
		score += stringsScore(dw.French, query)
		score += stringsScore(dw.English, query)
		return score
	}
}

func stringsScore(a string, b string) (score float64) {
	score = 0.0
	a = strings.ToLower(a)
	b = strings.ToLower(b)
	if a == b {
		score += 1.0
	} else {
		for i := 0; i < len(a)+1; i++ {
			for j := i + 2; j < len(a)+1; j++ {
				for k := 0; k < len(b)+1; k++ {
					for l := k + 2; l < len(b)+1; l++ {
						if a[i:j] == b[k:l] {
							score += 0.001 * math.Pow(float64(len(a[i:j])), 3)
						}
					}
				}
			}
		}
	}
	return
}