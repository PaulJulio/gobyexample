package main

import "fmt"

func main() {

	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// note that Println adds a space between arguments
	fmt.Println("a","b","c")

	fmt.Println("two\nlines")

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
