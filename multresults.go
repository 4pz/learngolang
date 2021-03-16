package main

import(
	"fmt"
)

func swap(x, y, z string) (string, string, string) {
	return y, z, x
}

func main() {
	a, b, c := swap("Hello", "world", "john")
	fmt.Println(a, b, c)
}