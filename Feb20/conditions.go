/*
Go Conditions:
- A condition can be either true or false.
- Go supports the usual comparison operators from mathematics
- Also supports the usual logical operators
*/


/*
Example:
x > y	
x != y	
(x > y) && (y > z)	
(x == y) || z
*/
package main
import ("fmt")

func main() {
  x := 10
  y := 5
  z := 2

  fmt.Println(x > y)
  fmt.Println(x != y)
  fmt.Println((x > y) && (y > z))
  // fmt.Println((x == y) || z) if z = false
}

