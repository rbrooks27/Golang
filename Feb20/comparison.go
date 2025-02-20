/*
Comparison Operators:
- Comparison operators are used to compare two values.
- The return value of a comparison is either true (1) or false (0).
*/

package main
import ("fmt")

func main() {

// Greater than operator (>)
  var x = 5
  var y = 3
  fmt.Println(x>y) // returns 1 (true) because 5 is greater than 3
  fmt.Println(x==y)// false (0)
  fmt.Println(x!=y) // true (1)
  fmt.Println(x<y) // false (0)
  fmt.Println(x>=y) // true (1)
  fmt.Println(x<=y) // false (0)
}

/*
A list of all comparison operators:

Operator			Name						Example	
	==				Equal to					x == y	
	!=				Not equal					x != y	
	>				Greater than				x > y	
	<				Less than					x < y	
	>=				Greater than or equal to	x >= y	
	<=				Less than or equal to		x <= y

*/