package httpService

import (
	"Dibi/src/dict"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func LaunchAPI() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/api", getDictionary)
	router.GET("/api/:query", getListByQuery)

	router.Run(Url + ":" + port)
}

// All words
func getDictionary(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dict.Dictionary)
}

// Words corresponding to a query
func getListByQuery(c *gin.Context) {
	query := c.Param("query")
	c.IndentedJSON(http.StatusOK, dict.GetMatchingList(query, 50))
}
