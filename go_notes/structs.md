July 9 2023

edited on July 10 2023

## Pointers in Go
1. Go pointers are used to reference the address of something. 
    a. This is typically used when passing `struct`'s by `refererence`
```go
i := 42
p = &i // get the pointer of i
*p = 21 // use * to de-reference the pointer
```
Go has *no* pointer arithmetic
1. You can't do things like increment a pointer or do `pointer comparisons`.
    this is where your brain typically starts to hurt when looking at low level
    c and c++ code
    a. In fact, Go was largely written due to the authors distain C++.
2. It is still not clear to me when you should be explicitly using pointers
in go. *Is it required for return types and params*?
    a. It appears now. See code below
    ```go
    package main

    import "fmt"

    type Vertex struct {
    	X int
    	Y int
    }
    func fuMuStruct(v Vertex) {
    	fmt.Println(v.X)
    }
    func main() {
    	v := Vertex{1, 2}
    	v.X = 32
    	fuMuStruct(v)
    }
    ```



## Structs
A collection of fields. Can use dot notation to access fields.
*`structs` are technically a pointer, but go abstracts the dereferncing
when using them. i.e. you can write `p.X` instead of `(*p).x`

```go
type Vertex struct {
	X int
	Y int
}
```

## Struct literal
The creation of a struct
```go
v1 = Vertex{1, 2} // where Vertext is a struct
```


