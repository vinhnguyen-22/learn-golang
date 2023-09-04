package main

import "fmt"

func main() {
 	fmt.Println("hello world") // print

	fmt.Println(sum(10)) // call function

	// type inference
	var email = "vincent@gmail.com"
	fmt.Println(email)

	// declare multiple variables
	var (
		a int
		b string
		c bool
	)

	a = 10
	b = "hello"
	c = true
	
	fmt.Println("multiple variables: ", a, b, c)

	// declare multiple variables with short declaration
	var d, e, f = 20, "hello", true
	fmt.Println("short declaration: ",d, e, f)

	// declare multiple variables with the same type
	var g, h, i int 
	g = 10
	h = 11
	i = 12
	fmt.Println("same type: ",g, h, i)

	// declare short-hand variable
	j := "hello"
	fmt.Println("short-hand variable: ", j)
}


func sum(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}