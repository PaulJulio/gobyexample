package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("empty:", a)
	a[4] = 100
	fmt.Println("set idx 4:", a)
	fmt.Println("get ids 4:", a[4])
	fmt.Println("length:", len(a))

	b := [5]int{1,2,3,4,5}
	fmt.Println("declared as:", b)

	const rows = 2
	const cols = 3
	var twoD [rows][cols]int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			twoD[i][j] = i * 10 + j
		}
	}
	fmt.Println("2d", twoD)
}