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

### Pointer
* ```*T``` is a pointer to a ```T```
```go
  var p *int
```
* ```&``` creates a pointer
```go
  i := 42
  p = &i
```
* ```*``` denotes the pointer's underlying value
```go
  fmt.Println(*p) // read i
  *p = 21 // set i with 21
```

### Arrays
* Declare
```go
  var a [2]string
  a[0] = "hello"
  a[1] = "world"
  fmt.Println(a[0], a[1])
  fmt.Println(a)

  primes := [6]int{2, 3, 5, 7, 11, 13}
```
* Slices
```go
  a := [6]int{2, 3, 5, 7, 11, 13}
  var s []int = a[1:4] // [3 5 7]
  b := a[0:3] // [2 3 5]
  c := a[:3] // [2 3 5]
  d := a[4:] // [7 11 13]
  e := [:] // [2 3 5 7 11 13]
  f := a[0:-1] // error

  q := []int{2, 3, 5, 7} // [2 3 5 7]
  s := []struct {
        i int
        b bool
  } {
        {2, true},
        {3, false}
  } // [{2 true} {3 false}]
```
* Zeror value of a slice is ```nil```
* ```make``` creates a slice with zero values
```go
  a := make([]int, 5) // [0 0 0 0 0]
  b := make([]int, 0, 4) // [], capacity is 5
```
* Multi-dimension
```go
  board := [][]string {
        []string{"_", "-"},
        []string{"-", "_"},
  }

  board[0][0] = "X"
  board[1][0] = "O"
```
* Append
```go
  var a []int
  a = append(a, 0) // [0]
```
* Range
```go
  var a = []int{1, 2, 4, 8}
  // i is index, v is value of index
  for i, v := range a {
    ...
  }
```

### Map
* Declare
```go
  type Point struct {
    x, y int
  }

  var m map[string]Point
  m1 = make(map[string]Point)
  m2 := map[string]Point {
        "p1" : Point { 0, 0 },
        "p2" : Point { 1, 1 },
  }
```
* Mutating Maps
```go
  m := make(map[string]int)
  m["x"] = 1 // [{"x" : 1}]
  m["x"] = 2 // [{"x" : 2}]
  m["y"] = 3 // [{"x" : 2}, {"y" : 3}]
  delete(m, "x") // [{"y" : 3}]
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
* Closure
```go
  // return function: func(int) int
  func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
  }
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

* Methods
```go
  func (p Person) Age() int {
    return p.age
  }

  a := Person{"Jack", 20}
  fmt.Println(a.Age()) // 20
```

* Method on Non-Struct
```go
  type MyFloat float64
  func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
  }

  f := MyFloat(-1.2)
  fmt.Println(f.Abs()) // 1.2
```
* Pointer Receivers
```go
  // age of p will be changed
  func (p *Person) setAge(age int) {
    p.age = age
  }

  func (p Person) Age() int {
    return p.age
  }
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
  else if y > 10 {
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
  // Print: hello world
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

### Interface
* Declare
```go
  type IPerson interface {
    Name() string
    Age() int
  }

  type Person struct {
    name string
    age int
  }

  type Man struct {
    name string
    age int
  }

  func (p *Person) Name() string {
    return p.name
  }

  func (p *Person) Age() int {
    return p.age
  }

  var p IPerson
  p.Name() // error since p is nil
  x := Person{"Jack", 20}
  p = x // Person implements IPerson

  y := Man{"Tom", 30}
  p = y // error
```
* Empty Interface
```go
  var i interface{}
  describe(i) // (<nil>, <nil>)

  i = 42
  describ(i) // (32, int)

  i = "hello"
  describe(i) // (hello, string)
```
* Type Assertions
```go
  var i interface{} = "hello"

  s := i.(string) // hello
  s, ok := i.(string) // hello true
  f, ok := i.(float64) // 0 false
  f = i.(float64) // panic
```
* Type Switches
```go
  func do(i interface{}) {
    switch v := i.(type) {
    case int:
        ...
    case string:
        ...
    default:
        ...
  }
```
* ```String()``` like ```toString()``` in java
```go
  func (p Person) String() string {
        return fmt.Printf("Name:%s, Age:%d", p.name, p.age)
  }
```

### Errors
* ```error``` is built-in interface
```go
  type error interface {
    Error() string
  }
```

* Example
```go
  type MyError struct {
    When time.Time
    What string
  }

  func (e *MyError) Error() string {
    return fmt.Sprintf("at %v, %s", e.When, e.What)
  }
```

### Goroutines: A lightweight thread maanged by Go runtime
* Run function in a thread
```go
  go f(x, y, z) // f(..) will be run in thread
```

### Channels
* Create Channel
```go
  c := make(chan int) // int is payload of channel
```

* Send value to channel
```go
  c <- v // send v to channel c
```

* Receive value from channel
```go
  v := <-c // receive from channel c
```
By default, sends and receives block until the other side is ready.

* Buffered Channels
```go
  ch := make(chan int, 100) // channel buffer length is 100
```

* Range and Close
```go
  func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)
  }

  func main() {
    c := make(chan int, 10)
    go fibonacci(cap(c), c)
    for i := range c { // range like foreach in other languages
        fmt.Println(i)
    }
  }
```

### sync.Mutex
```go
  type SafeCounter struct {
    v map[string]int
    mux sync.Mutex
  }

  func (c *SafeCounter) Inc(key string) {
    c.mux.Lock()
    c.v[key]++
    c.mux.Unlock()
  }

  func (c *SafeCounter) Value(key string) int {
    c.mux.Lock()
    defer c.mux.Unlock()
    return c.v[key]
  }

  func main() {
    c := SafeCounter{v : make(map[string]int)} // don't need to init mutex?
    for i := 0; i < 1000; i++ {
        go c.Inc("key1")
    }

    time.Sleep(time.Second)
    fmt.Println(c.Value("key1"))
  }
```
