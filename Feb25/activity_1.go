package main

import (
	"fmt"
)

func main() {
	// 1. Check if a number is positive
	num := 10
	if num > 0 {
		fmt.Println(num, "is positive")
	} else if num < 0 {
		fmt.Println(num, "is negative")
	} else {
		fmt.Println(num, "is zero")
	}

	// 2. Determine if the number is even or odd
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 3. Compare two numbers
	a, b := 15, 20
	if a > b {
		fmt.Println(a, "is greater than", b)
	} else if a < b {
		fmt.Println(a, "is less than", b)
	} else {
		fmt.Println(a, "is equal to", b)
	}

	// 4. Iterate through an array and print its elements
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Array elements:")
	for _, val := range arr {
		fmt.Println(val)
	}

	// 5. Perform basic arithmetic operations
	x, y := 12, 4
	fmt.Println("Addition:", x, "+", y, "=", x+y)
	fmt.Println("Subtraction:", x, "-", y, "=", x-y)
	fmt.Println("Multiplication:", x, "*", y, "=", x*y)
	fmt.Println("Division:", x, "/", y, "=", x/y)
}

/*
10 is positive
10 is even
15 is less than 20
Array elements:
1
2
3
4
5
Addition: 12 + 4 = 16
Subtraction: 12 - 4 = 8
Multiplication: 12 * 4 = 48
Division: 12 / 4 = 3
*/