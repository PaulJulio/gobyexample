package main

import "fmt"

func main() {
	fmt.Println("map.go")

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 11

	fmt.Println("map", m)

	v1, p1 := m["k1"]
	fmt.Println("value", v1, "set?", p1)

	v2, p2 := m["k3"]
	fmt.Println("value", v2, "set?", p2)

	_, p3 := m["k3"]
	fmt.Println( "set?", p3)

	// fmt.Println("blank", _) // the _ is the "blank identifier"

	n := map[int]string{7:"foo", 11:"bar"}
	fmt.Println("declared", n)
}