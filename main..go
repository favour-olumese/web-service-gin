package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// The json struct tags are used to dteremine what the
// field name would be called when serialized into JSON.
// Without them, the Title case of the names would be used.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// In memory data.
// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Onise Iyanu", Artist: "Nathaniel Bassey", Price: 200},
	{ID: "2", Title: "I Exalt You", Artist: "Kim Walker", Price: 230},
	{ID: "3", Title: "I Will Sing", Artist: "Don Moen", Price: 180},
	{ID: "4", Title: "Grace", Artist: "Michael Smith", Price: 220},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumID)
	router.POST("/albums", postAlbums)

	// Use the Run function to attach the router to an http
	router.Run("localhost:8080")
}

// gin.Contextcarries the request details, validates, and serializes JSON.
// Context.IndentedJSON serializes the struct into JSON and adds it to the
// You could also make use of Context.JSON, but this would make it compact.

// getAlbums reponds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// We use the curl commmand to make a request to our running web service
// Go does not enforce the order in which we declare functions.

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
	// StatusCreated implies 201 status code.
}

// Context.Param is used to retriee id path parameter from the URL.
// StatusNotFound - HTTP 404 error

// getAlbumID locates the album whos ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of album, looking for
	// an album whose ID value matches the parameter

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// TO GET FROM THE TERMINAL
/*
curl http://localhost:8080/albums
*/

// TO POST VIA THE TERMINAL
/*
curl localhost:8080/albums --include --header "Content-Type: application/json" --request "POST" --data "{\"id\": \"5\", \"title\": \"Be Magnified\", \"artist\": \"Don Moen\", \"price\": 178}"
*/
