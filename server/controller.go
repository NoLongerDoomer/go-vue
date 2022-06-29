package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type artist struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Bio  string `json:"bio"`
}

var artists = []artist{
	{ID: 1, Name: "Swarup", DOB: "14-12-1999", Bio: "something"},
	{ID: 1, Name: "Swarup", DOB: "14-12-1999", Bio: "something"},
}

func main() {
	router := gin.Default()
	router.GET("/getartists", getArtists)
	router.Run("localhost:6060")
}

func getArtists(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, artists)
}
