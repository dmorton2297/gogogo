// THIS CODE IS COPIED FROM THIS TUTORIAL https://go.dev/doc/tutorial/web-service-gin

package main

// Convention to put extra line between non standard imports
// and external imports
// Worth looking at `go.mod` to see how the dependencies were resolved
import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"` // specified the JSON key for serialization / deserialzation
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99}, // Defined just like JSON!
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99}, // All capital attributes 🤮
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
    router := gin.Default() // gin is a package to stnad up webservers

    // Stand up routing
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)

    // Run router on port
    router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    // indented JSON should only be used in development
    // This pretty formats the raw responsees with whitespace
    // and end line delimeters. This is recommended to be avoided
    // in prod as it puts extra cyles on the cpu per request
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    // What is gin context?
    // From docs: gin.Context is the most important part of Gin. It carries request details, validates and s    erializes JSON, and more
    // Gin context gets passed from the router
    id := c.Param("id")

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a) // Replace with c.JSON
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
