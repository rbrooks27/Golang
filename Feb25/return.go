/*
Function Returns
	- If you want the function to return a value, you need to define the data type of the return value (such as int, string, etc)
	- Also use the return keyword inside the function:

Syntax:
	func FunctionName(param1 type, param2 type) type {
	// code to be executed
	return output
	}
*/

// Examples

package main
import ("fmt")

// Example: Return
func myFunction(x int, y int) int {
  return x + y
}

// Named Return Values 
func myFunc(a int, b int) (result int) {
	result = a + b
	return
  }

// Storing the Return Value in a Variable
func myF(q int, p int)(result int){
	result = q + p
	return
}

// Multiple Return Values
func newFunction(k int, l string)(result int, txt1 string){
	result k + k
	txt1 = y + "Bong"
	return
}

// Storing the two return values into two variables
func double_it(f int, g string)(result int, txt2 string){
	result = f * f
	txt2 = g + "James"
	return
}

// Omitting the returned values
func omitting(t int, v string)(result int, txt3 string){
	result = t / t
	txt3 = v + "Boom, Boom"
	return
}

func main() {
  // Regular Return
  fmt.Println(myFunction(1, 2)) // Output: 3
  
  // Named Return Values 
  fmt.Println(myFunc(2, 3)) // Output: 5

  // Storing the Return Value in a Variable
  total_of_myF := myFunc(6, 4)
  fmt.Println(total_of_myF) // Output: 10

  // Multiple Return Values
  fmt.Println(newFunction(7, "Bing")) // Output: 14 Bing Bong

  // Storing the two return values into two variables
  e, r := double_it(6, "Lebron")
  fmt.Println(e, r) // Output: 36 Lebron James

  // Omitting the first returned value 
  _, v := omitting(7, "Boom,")
  fmt.Println(v) // Output: Boom, Boom, Boom 

  // Omitting the first returned value 
  t, _ := omitting(7, "Boom,")
  fmt.Println(v) // Output: 1

} 



