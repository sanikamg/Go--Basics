package main

import "fmt"

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, b)
}
