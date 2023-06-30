Based on https://go.dev/tour/list

Intention: Gain some more familiarity with Go



## Things from yesterday
1. := is declaration and assignment
    a. `number := 10` === `var number int = 10`
    b. The technical term fo this in go is `short assignment`
2. Go code is organized into `modules`
3. entry points in go are the func main() in some .go file within a module
4. To generate a new go module, run `go mod init <domain>/<module_name> 
    a. Under the hood, this configures a `go.mod` file that declares a folder as a go module 
    b. go build will create an executable for you to run
5. Tests are defined with `something_test.go` and test cases are functions that start with `Test`.
    a. Furthermore, by `import "testing"` at the top of the go file, you gain access to test helpers

## Basic go types
1. ```go
    bool

    string

    int  int8  int16  int32  int64
    uint uint8 uint16 uint32 uint64 uintptr

    byte // alias for uint8

    rune // alias for int32
     // represents a Unicode code point

    float32 float64

    complex64 complex128
    ```
2.  type conversion example: 
    ```go
    var i int = 42
    var f float64 = float64(i)
    ```
3. type inference is a thing

## Declaring arrays
a. an array type is specified as `[]type` or `[n]type` where `n` is the number of entries
in a known length array.

## For loop 
1. for loop syntax
```go
for i := 0; i < 10; i ++ {}
```
2. "While" loop syntax
```go
for some_bool_condition {} // COOL!
```

## Nil vs Zero values
1. `0 | false | ""` given to variables without explicit initial value

## Variables
1. `var <list of arguments> type`
2. can initialize i.e. `var i, j int = 1, 2`
3. short var declaration := (can only be used within functions)
4. Constants can be declared using the `const` keyword. They cannot be
declared using the `:=` syntax.

## Functions
1. Function syntax
    ```go
    func add(x int, y int) int {
	    return x + y
    }
    ```
2. If you have consecutive variables that share types, you can short hand
    a. `x int, y int === x, y int`
3. You can return a number of return results from a function. 
    ```go
    package main

    import "fmt"

    func swap(x, y string) (string, string) {
	    return y, x
    }

    func main() {
	    a, b := swap("hello", "world")
	    fmt.Println(a, b)
    }
    ```
4. Named return values allow you to instantate variables at the top of functions and specify
what they return. See more https://go.dev/tour/basics/7

## Packages / Executables
1. `package main` at the top of go executables
    a. This differentiates to the compiler that what you are trying to build is a go executable
    and not a package
    b. To create a go `module` that won't be interpreted as a package, use `package <module_name>`
    at the top of a `.go` source file
2. Package names are always the last word of the import statements
    a. for example `import "math/rand` imports the `rand` package into your go source.
3. Grouped import statements are called `factored` import statements. This is preferred
    ```go
        // Notice no commas or semicolons, weird!
        import (
	        "fmt"
	        "math"
        )
    ```


Excercise today: Wrote a bubble sort
