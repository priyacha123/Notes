package main

import "fmt"
import "math/rand"

func main() {
	var i int
	i = 10
	fmt.Println(i)
	fmt.Println("Random number", rand.Intn(100)) // "rand" gives a random number between 0 and 99
} 