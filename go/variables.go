package main

import "fmt"

func main() {
	var x int
	x = 1
	fmt.Printf("x = %d\n", x)

	var y int = 2
	fmt.Printf("y = %d\n", y)

	z := 3
	fmt.Printf("z = %d\n", z)

	i, z := 4, 4
	fmt.Printf("z = %d, i = %d\n", i, z)
}
