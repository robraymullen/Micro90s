package rest

import (
	"guestbook/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	router.GET("/guests", GetGuests)
	router.Run("localhost:8080")
}

func GetGuests(c *gin.Context) {
	guests := db.RunDB()
	c.IndentedJSON(http.StatusOK, guests)
}
