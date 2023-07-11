## Slices
Describes an underlying array


## Slicing out of slices
```go
package main

import "fmt"

func main() {
    // Array because it has a length provided
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// Slice - Reference to the memory in primes
    var s []int = primes[1:4]
	fmt.Println(s)
}

```
1. Slice literal is like declaring an array without a length (s variable above)
2. A nil slice is when there is no reference to an underlying array. Essentially
a nil pointer
