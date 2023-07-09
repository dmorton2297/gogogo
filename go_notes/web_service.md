Jul 9 2023

## What I did last time
1. Worked on a Project Euler problem to gain more familiarity with the language.
2. Where I need some work - Slices (array) and runes
3. if (var assignment); expression {} is nice
4. Had some exposure to `strconv` standard lib. Can be used for formatting differnt
things as strings.
5. Handling errors as values is nice


## What I am doing today
1. Standing up a simple web server based on this guide: https://go.dev/doc/tutorial/web-service-gin
2. Based on the gin web framework https://gin-gonic.com/docs/

## Gin
Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API, but with performance up to 40 times faster than Martini. If you need smashing performance, get yourself some Gin.
1. The fuck is a martini like API?
    a. Martini is a package for writing modular web application/services in Golang

## Building API's with Gin
1. The setup for Gin reminds me alot of the how a node express server works. In the main function of the go program, you initialize a Gin router, then define your routes with the relavent handlers.

See webservice Go module in `go_pf` directory


```go
func main() {
    router := gin.Default() // gin is a package to stnad up webservers

    // Stand up routing
    router.GET("/stuff", getStuff)
    router.POST("/stuff", postStuff)

    // Run router on port
    router.Run("localhost:8080")
}

// Passes in the gin context pointer
func getStuff(c *gin.Context) {
    // do stuff
}
```


## What I need to work more on
1. Pointers in Go 
