Created July 10 2023

Largely based off of information found here: https://go.dev/tour/moretypes/20

## Maps
1. Maps are Go's implementation of a dict
2. They contain `key` `value` pairs
Example usage
```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

```
1. Map literals are the instantiation of a map using literals
```go
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
```
2.Delete an element from map
```go
delete(m, key)
```
3. Check for presence in map
```go
elem, ok = m[key] // ok will be a boolean
```
