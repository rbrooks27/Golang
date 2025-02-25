// Factorial Recursion

/*
What is a factorial 
- n! = n * (n-1)* (n-2)
*/

package main
import ("fmt")

func factorial_recursion(x float64) (y float64) {
  if x > 0 {
     y = x * factorial_recursion(x-1)
  } else {
     y = 1
  } // base case
  return
}

func main() {
  fmt.Println(factorial_recursion(4))
} // Output: 24