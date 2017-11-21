package main

import "fmt"

func main() {
	fmt.Println("range.go")

	nums := []int{2,3,4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum", sum)

	for i, num := range nums {
		fmt.Println(i, "->", num)
	}

	kvs := map[string]string{"foo":"bar", "bizz":"buzz"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	counter := 0
	for k := range kvs {
		fmt.Println("key", counter, "is", k)
		counter++
	}

	for i, c := range "PaulJulio" {
		fmt.Println(i,c) // Unicode code points
	}
}