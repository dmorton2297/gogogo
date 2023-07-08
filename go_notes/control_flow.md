Created: July 7 2023
Intention: More lanugage familiarity

## Review of what I did last time
1. The last time I picked looked into some core elements of the language
2. For loop syntax `for i := 0; i < 10; i ++ {}`
    a. we are implicitly declaring and initializing i
    b. i implicitly is a `var` in Go
    c. `i := 1 == var i int = 1`
    d. `const i int = 1`
3. Short declaration can only be used within functions

4. 👇Sick!
```go
package main

import "fmt"

// this reminds me alot of swift tuples
func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println(a, b)
}
```
5. To indicate that your go module is a package, use `package <module_name` at the top of your .go source. Anything that is not an executable is a package.
6. Array declaration `[]int{2, 6, 1, 2, 7, 4}`. Weird!


-------
Begin new notes


## for loop review
```go
// init statement; condition; post statement
// docs specifically make note that they don't
// user parens lol.
for i := 0; i < 10; i++ { ... }
```
1. Init and post statements are optional (go's implementation of while)
```go
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```
2. Infinite loop!
```go
for {...}

```

## If statements with 'short' conditions
```go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
```
1. If you else this statement - you can still access v. Reminds me alot of swift.
2. Newtons method - neat! If you are trying to guess a root - you can do
    a. x1 = x2 = (x^2 - c)/(2x)
    b. denominator is the derivative of the numerator.
    c. If you have any function that approahces 0, you can use this
    to iteratively make guesses until the numertoar becomes so small 
    you find a solution


## Switch
1. Biggest thing of note - now break statements 🤯
```go
switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
}
```
```go
// Equiv to a switch true { ... }
switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	default:
		fmt.Println("Good evening.")
	}
```


## Defer
Straight from the site: A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

```go
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

```

* Defers can be stacked, they are executed in a last-in-first-out order when the
surrounding function returns


















