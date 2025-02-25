/*
For Loops:
	- The for loop loops through a block of code a specified number of times.
	- The for loop is the only loop available in Go.
	- Loops are handy if you want to run the same code over and over again, each time with a different value. 
	- Each execution of a loop is called an iteration.
*/

/*
Syntax:
	for statement1; statement2; statement3 {
	// code to be executed for each iteration
	}

		- statement1 Initializes the loop counter value.
		- statement2 Evaluated for each loop iteration. If it evaluates to TRUE, the loop continues. If it evaluates to FALSE, the loop ends.
		- statement3 Increases the loop counter value
*/

// Example 
	package main
	import ("fmt")

	func main() {
	for i:=0; i < 5; i++ {
		fmt.Println(i)
	}
	/*
	i := 0; // intializes it to 0
	i < 5; // says it will keep running till it hits 5
	i++ // keey adding one to keep running so... 1+0 then 1+1 then 2+1 and so on.
	*/

	/* Output
			0
			1
			2
			3
			4
	*/
	}

/* Example 2

	package main
	import ("fmt")

	func main() {
	for i:=0; i <= 100; i+=10 {
		fmt.Println(i)
	}
	}

	Output: 
		0
		10
		20
		30
		40
		50
		60
		70
		80
		90
		100
*/



