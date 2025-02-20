// Activity
package main

import ("fmt")

func main() {
	num1 := 10.0
	num2 := 5.0
	operator := "+" 

	// Switch case 
	switch operator {
	case "+":
		fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, num1+num2)
	case "-":
		fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, num1-num2)
	case "*":
		fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, num1/num2)
		} else {
			fmt.Println("Error: Division by zero is not allowed.")
		}
	default:
		fmt.Println("Invalid operator. Please use +, -, *, or /.")
	}
}
