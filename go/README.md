### Main Function

```go
  func main(argc int, argv []string) int
  func main(int, []string) int
  func main()
```

### Variables

* Declare
```go
  var x int
  x = 1

  var y int = 1
  var c, python, java bool

  // short declaration
  z := 1

  i, j := 1
```

* Pointer
```go
  var a []int
  x = a[1]

  var p *int
  x = *p

  var p ^int
  x = p^
```

* Variable will be assigned default value after declared
  * **0** for numeric types
  * **false** for boolean type
  * **""** for strings

* Constant
```go
  const Pi = 3.14
```
  * Can't use **:=** syntax

### Basic Types

* Basic Types
  * bool
  * string
  * int, int8, int16, int32, int64
  * uint, uint8, uint16, uint32, uint64 uintptr
  * byte == unit8
  * rune == int32
  * float32, float64
  * complex64, complex128
* Type Conversions: ```T(v)```
```go
  var i int = 42
  var f float64 = float64(i)
  var u uint = uint(f)

  i := 42
  f := float64(i)
  u := uint(f)
```
* Type Inference
```go
  var i int
  j := i // j is an int

  i := 42 // int
  f := 3.142 // float64
  g := 0.867 + 0.5i // complex128
```

### Function

* Declare
```go
  // no return value
  func log(message string) {
  }

  // return one value
  func add(a int, b int) int {
  }

  // return two value
  func power(name string) (int, bool) {
  }

  // concise version
  func add(a, b int) int {
  }

  // skip one of return value
  x, _ := power("xx")
```

### Struct

* Declare
```go
  type Person struct {
    Name string
    Age int
  }
```

* Initialize
```go
  jack := Person {
    Name: "Jack",
    Age: 60
  }

  anonymous := Person {}

  tom := Person { Name: "Tom" }
  tom.age = 20

  adm := Person { "Adm", 30 }
```

### Pass value or pointer
```go
  // pass value
  func AddAge1(Person p) {
    p.Age += 10
  }

  jack := Person {"Jack", 50}
  addAge1(jack)

  // pass pointer
  func AddAge2(Person *p) {
    p.Age += 10
  }

  tom := &Person {"Tom", 50}
  addAge2(tom)
```

### Package
* Import
```go
  import "fmt"
  import (
    "fmt"
    "math"
  )
```
* Exported name begins with capital letter

### Controller

* **For**
```go
  for i := 0; i < 10; i++ {
    ...
  }

  sum := 1
  for ; sum < 1000; {
    sum += sum
  }

  // same as while
  sum := 1
  for sum < 1000 {
    sum += sum
  }

  // forever
  for {
  }
```

* **If**
```go
  if x < 0 {
  }

  if v := math.Pow(x, n); v < lim {
  }

  if y < 0 {
  }
  else {
  }
```

* **Switch**
```go
  switch os := runtime.GOOS; os {
  case "darwin":
    ...
  case "linux":
    ...
  default:
    ...
  }

  // evaluate cases from top to bootom,
  // stopping when a case succeeds
  switch i {
  case 0:
  case f():
  }

  // switch with no condition
  t := time.Now()
  switch {
  case t.Hour() < 12:
    ...
  case t.Hour() < 17:
    ...
  default:
    ...
  }
```

* **Defer**
  * A defer statement defers the execution of a function until the surrounding
function returns
```go
  // output: hello world
  func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
  }
```
  * **Stacking defers**, last-in-first-out order
```go
  func main() {
    fm.Println("couting")

    // Print: 9, 8, 7, 6 ... 0
    for i := 0; i < 10; i++ {
      defer fmt.Println(i)
    }

    fmt.Println("don")
  }
```
