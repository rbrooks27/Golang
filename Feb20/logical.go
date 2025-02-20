/*
Logical Operators
- Logical operators are used to determine the logic between variables or values

	Operator		Name								Description								Example	
	&& 				Logical and		Returns true if both statements are true					x < 5 &&  x < 10	
	|| 				Logical or		Returns true if one of the statements is true				x < 5 || x < 4	
	!				Logical not		Reverse the result, returns false if the result is true		!(x < 5 && x < 10)
*/

package main
import ("fmt")
func main() {

  var x = 5

  fmt.Println(x < 5 &&  x < 10) // false

  fmt.Println(x < 5 || x < 4) // false

  fmt.Println(!(x < 5 && x < 10)) // true

}