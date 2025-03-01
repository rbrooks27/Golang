/*
Functions:
	- A function is a block of statements that can be used repeatedly in a program.
	- A function will not execute automatically when a page loads.
	- A function will be executed by a call to the function
*/

/*
Creating Function/Syntax:
		- To create (often referred to as declare) a function, do the following:
				- Use the func keyword.
				- Specify a name for the function, followed by parentheses ().
				- Finally, add code that defines what the function should do, inside curly braces {}.

Syntax:				
	func FunctionName() {
	// code to be executed
	}
*/

/*
Calling a Function:
	- Functions are not executed immediately.
	- Call it in main with <functionname>()
*/

// Example
package main
import ("fmt")

func myMessage() {
  fmt.Println("I just got executed!")
} // Output: I just got called

func main() {
  myMessage() // call the function
  // myMessage() // this would call it again
}