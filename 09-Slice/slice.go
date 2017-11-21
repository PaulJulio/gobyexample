package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("empty:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	// s[3] = "error: index out of range"
	fmt.Println("set:", s)
	fmt.Println("get 2:", s[2])
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("appended", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy", c)

	l := s[2:5]
	fmt.Println("subset 2:5", l)

	l = s[2:]
	fmt.Println("subset 2:", l)

	l = s[:5]
	fmt.Println("subset :5", l)

	const rows = 3
	twoD := make([][]int, rows)
	for i := 0; i < rows; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i * 10 + j
		}
	}
	fmt.Println("2d:", twoD)
}
