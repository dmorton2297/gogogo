created June 28 2023

Based on documentation here: https://go.dev/doc/tutorial/create-module


## Dependency management
0. To create a new module, run `go mod init <domain>/<module_name>`
1. Go manages dependency manageament through `modules` that each have a `go.mod` file that specifies all external deps.
2. To find packages that you can depend on, visit https://pkg.go.dev/
3. To automatially update your deps in go.mod, import the code in your go file and run `go mod tidy` in the terminal.
4. To import a local go file, run `go mod edit -replace <name_imported>=<relative_or_absolute_path>`



## Writing tests
0. Write test files in a`<module_name>_test.go`
1. Test functions must begin with the keywork  `Test`
1. Use `import "testing"` to gain access to the following pattern
```go
// * referes to a pointer in this case
func TestCase(t *testing.T) {
    t.Fatalf('Some failure case if a condition fails')
}
```


## Building a go module
0. Run `go build` in the directory you are trying to build.
1. Can determine the install path of any module by running `$ go list -f '{{.Target}}'`
