package httpService

import (
	"Dibi/src/dict"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LaunchAPI() {
	router := gin.Default()

	router.GET("/api", getDictionary)
	router.GET("/api/:query", getListByQuery)

	router.Run(Url + ":" + APIPort)
}

// All words
func getDictionary(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dict.Dictionary)
}

// Words corresponding to a query
func getListByQuery(c *gin.Context) {
	query := c.Param("query")
	fmt.Printf("%s\n", query)
}
