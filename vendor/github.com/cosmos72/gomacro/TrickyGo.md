A collection of tricky go code

```
// change the meaning of true
const true = false
println(true)
```

```
// change the meaning of uint
type uint int
println(uint(1))
```

```
// change the meaning of uint (again)
func uint(x int) int { return x + 7 }
println(uint(1))
```

```
// nil interface values don't implement interfaces
var x error
y := x.(error) // panic!
```

```
// except that nil interface{} implements interface{}
var x interface{}
y := x.(interface{}) // works
```

```
import "os"
func getGoPath() string {
	dir := os.Getenv("GOPATH")
	if len(dir) == 0 {
		dir := os.Getenv("HOME") // shadows outer "dir", does NOT modify it
		if len(dir) == 0 {
			panic("cannot determine go source directory: both $GOPATH and $HOME are unset or empty")
		}
		dir += "/go"
	}
	return dir // inner "dir" is not seen -> always returns os.Getenv("GOPATH")
}
```
