package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{
		ID: "1", Title: "Hotel California", Artist: "Eagles", Price: 25.00,
	},
	{
		ID: "2", Title: "Happy Together", Artist: "The Turtles", Price: 32.25,
	},
	{
		ID: "3", Title: "What A Wonderful World", Artist: "Lance Armstrong", Price: 23.75,
	},
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	err := router.Run("localhost:8080")
	if err != nil {
		println("There is an error", err)
	}

}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, newAlbum)
}
