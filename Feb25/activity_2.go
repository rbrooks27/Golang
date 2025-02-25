package main

import (
	"fmt"
)

// 1. Check if a number is positive, negative, or zero
func checkNumberSign(num int) {
	if num > 0 {
		fmt.Println(num, "is positive")
	} else if num < 0 {
		fmt.Println(num, "is negative")
	} else {
		fmt.Println(num, "is zero")
	}
}

// 2. Determine if a number is even or odd
func checkEvenOdd(num int) {
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}

// 3. Compare two numbers
func compareNumbers(a, b int) {
	if a > b {
		fmt.Println(a, "is greater than", b)
	} else if a < b {
		fmt.Println(a, "is less than", b)
	} else {
		fmt.Println(a, "is equal to", b)
	}
}

// 4. Iterate through an array and print its elements
func printArrayElements(arr []int) {
	fmt.Println("Array elements:")
	for _, val := range arr {
		fmt.Println(val)
	}
}

// 5. Perform basic arithmetic operations
func performArithmeticOperations(x, y int) {
	fmt.Println("Addition:", x, "+", y, "=", x+y)
	fmt.Println("Subtraction:", x, "-", y, "=", x-y)
	fmt.Println("Multiplication:", x, "*", y, "=", x*y)
	if y != 0 {
		fmt.Println("Division:", x, "/", y, "=", x/y)
	} else {
		fmt.Println("Division by zero is not allowed.")
	}
}

func main() {
	num := 10
	checkNumberSign(num)
	checkEvenOdd(num)

	a, b := 15, 20
	compareNumbers(a, b)

	arr := []int{1, 2, 3, 4, 5}
	printArrayElements(arr)

	x, y := 12, 4
	performArithmeticOperations(x, y)
}
