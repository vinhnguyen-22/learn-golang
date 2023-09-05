package main

import (
	"fmt"
)

/**
Primitive data types

1. boolean
2. Numeric data types
	1. int
		signed int : int8 int16 int32 int64 (int8 = 8 bit)
		unsigned int : uint8 uint16 uint32 uint64  cannot store negative number
	2. float float32 float64
	3. complex complex64 complex128
3. Text
	1. string
	2. byte -> UTF-8
	3. rune -> UTF-32

*/

func main() {

	var (
		i int    = 10
		j string = "Hello world"
		k bool   = false

		a int8 = -10
		b int8 = 20

		c int8 = 10
		d int8 = 3
	)

	fmt.Printf("Value of i: %v, a: %v,b: %v, j: %v, k: %v\n", i, a, b, j, k)
	fmt.Printf("Type of i: %T, a: %T,b: %T, j: %T, k: %T\n", i, a, b, j, k)

	fmt.Print("===== operator =====\n")
	fmt.Println("operator: ", operator(a, b, "+"))
	fmt.Println("operator: ", operator(a, b, "-"))
	fmt.Println("operator: ", operator(a, b, "*"))
	fmt.Println("operator: ", operator(a, b, "/"))
	fmt.Println("operator: ", operator(a, b, "%"))

	fmt.Print("===== operator with bit =====\n")
	fmt.Println("operatorWithBit: ", operatorWithBit(c, d, "|"))
	fmt.Println("operatorWithBit: ", operatorWithBit(c, d, "^"))
	fmt.Println("operatorWithBit: ", operatorWithBit(c, d, "&"))
	fmt.Println("operatorWithBit: ", operatorWithBit(c, d, "&^"))

}

func operator(a int8, b int8, operator string) any {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	case "%":
		return a / b
	}
	return 0
}
func operatorWithBit(a int8, b int8, operator string) any {
	switch operator {
	case "|":
		return a + b
	case "&":
		return a - b
	case "^":
		return a * b
	case "&^":
		return a / b
	}
	return 0
}
