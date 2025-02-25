/*
The switch Statement:
	- Use the switch statement to select one of many code blocks to be executed.
	- The switch statement in Go is similar to the ones in C, C++, Java, JavaScript, and PHP. The difference is that it only runs the matched case so it does not need a break statement.


*/

/*
Syntax:

	switch expression {
	case x:
	// code block
	case y:
	// code block
	case z:
	...
	default:
	// code block
	}
*/

// Example of a Single-Case switch
package main
import ("fmt")

func main() {
  day := 4

  switch day {
  case 1:
    fmt.Println("Monday")
  case 2:
    fmt.Println("Tuesday")
  case 3:
    fmt.Println("Wednesday")
  case 4:
    fmt.Println("Thursday")
  case 5:
    fmt.Println("Friday")
  case 6:
    fmt.Println("Saturday")
  case 7:
    fmt.Println("Sunday")
  }
}


// Output: Thursday

/*
Multi-case switch Statement:
	- It is possible to have multiple values for each case in the switch statement

Syntax:
	switch expression {
	case x,y:
	// code block if expression is evaluated to x or y
	case v,w:
	// code block if expression is evaluated to v or w
	case z:
	...
	default:
	// code block if expression is not found in any cases
	}
*/

/* Exmple:

package main
import ("fmt")

func main() {
   day := 5

   switch day {
   case 1,3,5:
    fmt.Println("Odd weekday")
   case 2,4:
     fmt.Println("Even weekday")
   case 6,7:
    fmt.Println("Weekend")
  default:
    fmt.Println("Invalid day of day number")
  }
}
 // Output: Odd weekday

*/
