/*
Recursion Functions:
	- Go accepts recursion functions. 
	- A function is recursive if it calls itself and reaches a stop condition.

*/

// Example

package main
import ("fmt")

func testcount(x int) int {
  if x == 11 {
    return 0
  } // Base Case
  fmt.Println(x)
  return testcount(x + 1) // recursive call
}

func main(){
  testcount(1)
}