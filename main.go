package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{
		Id: "1", Title: "Hotel California", Artist: "Eagles", Price: 25.00,
	},
	{
		Id: "2", Title: "Happy Together", Artist: "The Turtles", Price: 32.25,
	},
	{
		Id: "3", Title: "What A Wonderful World", Artist: "Lance Armstrong", Price: 23.75,
	},
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)

	err := router.Run("localhost:8080")
	if err != nil {
		println("There is an error", err)
	}

}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	for _, alb := range albums {
		if alb.Id == id {
			c.IndentedJSON(http.StatusOK, alb)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, nil)
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
