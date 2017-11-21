package main

import "fmt"
import "math/rand"
import "time"

func main() {
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(20)
	b := rand.Intn(52)

	if a%2 == 0 {
		fmt.Printf("%d is even\n", a)
	} else {
		fmt.Printf("%d is odd\n", a)
	}

	if b%4 != 0 {
		fmt.Printf("%d is not divisible by 4\n", b)
	}

	if c := rand.Intn(20) - 10; c < 0 {
		fmt.Printf("%d is negative\n", c)
	} else if c < 10 {
		fmt.Printf("%d has 1 digit\n", c)
	} else {
		fmt.Printf("%d has multiple digits\n", c)
	}
}