### Writing tests

Writing a test is just like writing a function, with a few rules

* It needs to be in a file with a name like `xxx_test.go`
* The test function must start with the word `Test`
* The test function takes one argument only `t *testing.T`
* In order to use the `*testing.T` type, you need to `import "testing"`, like we did with `fmt` in the other file

### Go doc

Godoc install it with `go get golang.org/x/tools/cmd/godoc`

You can launch the docs locally by running `godoc -http :8000`. If you go to [localhost:8000/pkg](http://localhost:8000/pkg) you will see all the packages installed on your system.

`go test -v`

### Benchmarking

Writing [benchmarks](https://golang.org/pkg/testing/#hdr-Benchmarks) in Go is another first-class feature of the language and it is very similar to writing tests.

```go
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
```

You'll see the code is very similar to a test.

The `testing.B` gives you access to the cryptically named `b.N`.

When the benchmark code is executed, it runs `b.N` times and measures how long it takes.

The amount of times the code is run shouldn't matter to you, the framework will determine what is a "good" value for that to let you have some decent results.

To run the benchmarks do `go test -bench=.` (or if you're in Windows Powershell `go test -bench="."`)

```text
goos: windows
goarch: amd64
pkg: github.com/mhmmdd/go-with-tests/03-iterations
cpu: Intel(R) Core(TM) i7-6600U CPU @ 2.60GHz
BenchmarkRepeat-4        8733344               131.5 ns/op
PASS
ok      github.com/mhmmdd/go-with-tests/03-iterations   1.513s
```

#### Finding unchecked errors
There is one scenario we have not tested. To find it, run the following in a terminal to install `errcheck`, one of many linters available for Go.

`go get -u github.com/kisielk/errcheck`

Then, inside the directory with your code run `errcheck .`

You should get something like

`wallet_test.go:17:18: wallet.Withdraw(Bitcoin(10))`

### Map value is a pointer

An interesting property of maps is that you can modify them without passing as an address to it (e.g `&myMap`)

This may make them _feel_ like a "reference type", [but as Dave Cheney describes](https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it) they are not.

> A map value is a pointer to a runtime.hmap structure.

So when you pass a map to a function/method, you are indeed copying it, but just the pointer part, not the underlying data structure that contains the data.

A gotcha with maps is that they can be a `nil` value. A `nil` map behaves like an empty map when reading, but attempts to write to a `nil` map will cause a runtime panic. You can read more about maps [here](https://blog.golang.org/go-maps-in-action).

Therefore, you should never initialize an empty map variable:

```go
var m map[string]string
```

Instead, you can initialize an empty map like we were doing above, or use the `make` keyword to create a map for you:

```go
var dictionary = map[string]string{}

// OR

var dictionary = make(map[string]string)
```

Both approaches create an empty `hash map` and point `dictionary` at it. Which ensures that you will never get a runtime panic.