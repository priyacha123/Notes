package main

import "fmt"
// import "math/rand"

// func main() {
// 	var i int
// 	i = 10
// 	fmt.Println(i)
// 	fmt.Println("Random number", rand.Intn(100)) // "rand" gives a random number between 0 and 99
// } 




// functions

// it is implicit returning of values
func swap(x, y string) (string, string) {
	return y, x
}

// here we are returning two named values, so we can just return without specifying the names of the return values
// it is explicit returning of values
func split(sum int) (x,y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// := is a short variable declaration, it can only be used inside functions
func main() {
	a, b := swap("5", "10")
	fmt.Println(a, b)
	fmt.Println(split(17))
}