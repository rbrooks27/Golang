// Introduction
	/* What is Go?
		- Golang is a open source, fast, statically typed, compiled programming language
		- Golang is a procedural programming language  
		- Known for it fast complie time and automatic garbage collection
		- Golang does not support inheritance or classes/objects
		- Golang is case sensitive
	*/

// Go Syntax
	/*
	- Every Go programs need the following things to run:
		1. package declaration (package main)
		2. import packages (import ("fmt"))
		3. Functions - you need a main (func main(){})
		4. Statements and expression - the code you want to make

	- For print statement the following syntax must be used:
			fmt.Println("")

	*/

// Go Comments 
	/*
	- A comment is a text that is ignored upon execution.
	- Great to use to explain your code or make it more readable

	Use: // for single-line comments
	Use: / * * / for multiple comments
	*/

// Go Variables 
	/*

	Variable Types:
		- int  (default value = 0)
		- float32 (default value = 0.0)
		- string (default value = "")
		- bool (default value = false)

	Declaring Variable Syntax without initial value:
		var <variablename> type

		Ex. 
		var Letter_Grade string = "A+"

	Declaring Variable Syntax with initial value:

		var <variablename> type = value
					or
		var <variablename> := value         (this is called inferring where the compiler assume the type based on the value)
											note this can only be done inside functions 

		Ex.
		var Lebron_Number int = 23
		var Starbucks_cost := 12.34  // inferring float32
		*/

// Multiple Variable Declarartion:
	/*
	Syntax: 
		var <vari1>, <vari2>, <vari3>, <vari4> type = <value1>, <value2>, <value3>, <value4>
													or
		var (
		<vari1> = <value1>
		<vari2> = <value2>
		<vari3> = <value3>
		<vari4> = <value4>
		)

	*/
// Variable Naming Rules:
 /*
		Naming Rules:
			- A variable name must start with a letter or an underscore character (_)
			- A variable name cannot start with a digit
			- Variable names are case-sensitive 
			- The variable name cannot be any Go keywords (like range)
		
		Multi-Word Variable Names:
			Camel Case: myVariableName = "John"
			Pascal Case: MyVariableName = "John"
			Snake Case: my_variable_name = "John"
		

	
 */

// Constants: 
 /*
	- Use the keyword: const to declare a constant
	- Syntax:
		const CONSTNAME type = value
	- Two types of constants
		1. Typed: const A int = 1
		2. Untyped: const A = 1    (inferring again)
	- You can do a block declaration too
		const(
		<variablename> type = value
		<variablename> type = value
		<variablename> type = value
		)

 */
// Output Functions: 
 /*
 Print()
	- function prints its arguments with their default format (so no spaces but \n can fix it)
	- Ex. HelloWorld 
 Println()
	- similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end
	- Ex. Hello World
 Printf()
	- first formats its argument based on the given formatting verb and then prints them.
	- two formatting verbs:
		1. %v is used to print the value of the arguments
		2. %T is used to print the type of the arguments
	- Ex. 
		fmt.Printf("i has value: %v and type: %T\n", i, i) // Output: i has value: Hello and type: string
  		fmt.Printf("j has value: %v and type: %T", j, j)   // Output: j has value: World and type: string

 Formatting Verbs:
 	
 	General Formatting:
			Verb	Description
			%v		Prints the value in the default format
			%#v		Prints the value in Go-syntax format
			%T		Prints the type of the value
			%%		Prints the % sign
	Interger Formatting:
			Verb	Description
			%b		Base 2
			%d		Base 10
			%+d		Base 10 and always show sign
			%o		Base 8
			%O		Base 8, with leading 0o
			%x		Base 16, lowercase
			%X		Base 16, uppercase
			%#x		Base 16, with leading 0x
			%4d		Pad with spaces (width 4, right justified)
			%-4d	Pad with spaces (width 4, left justified)
			%04d	Pad with zeroes (width 4)
	String Formatting Verbs:
			Verb	Description
			%s		Prints the value as plain string
			%q		Prints the value as a double-quoted string
			%8s		Prints the value as plain string (width 8, right justified)
			%-8s	Prints the value as plain string (width 8, left justified)
			%x		Prints the value as hex dump of byte values
			% x		Prints the value as hex dump with spaces
	Boolean Formatting Verbs;
			Verb	Description
			%t		Value of the boolean operator in true or false format (same as using %v)

	Float Formatting Verbs
			Verb	Description
			%e		Scientific notation with 'e' as exponent
			%f		Decimal point, no exponent
			%.2f	Default width, precision 2
			%6.2f	Width 6, precision 2
			%g		Exponent as needed, only necessary digits
 */
// Data Types
 /*
 - Go is statically typed

 Basic Data Types:
  - bool: represents a boolean value and is either true or false
		• The default value of a boolean data type is false.
  - Numeric (int): represents integer types, floating point values, and complex types
		• Signed integers - can store both positive and negative values
		• Unsigned integers - can only store non-negative values
  - string: represents a string value
		• used to store a sequence of characters (text). String values must be surrounded by double quotes
		• Ex. var txt1 string = "Hello!"
  - float: represent decimals 
		• type: float32	size: 32 bits	range: -3.4e+38 to 3.4e+38
		• type: float64	size: 64 bits	range: -1.7e+308 to +1.7e+308.
 */
// Arrays:
 /*
 Declaring an Array:
	 Syntax:
	 						With Keyword Var:
	 	var array_name = [length]datatype{values} // here length is defined
					   or
		var array_name = [...]datatype{values} // here length is inferred

							Without Keyword Var:
		array_name := [length]datatype{values} // here length is defined
						or
		array_name := [...]datatype{values} // here length is inferred

	Example:
		var arr1 = [3]int{1,2,3} //
		arr2 := [5]int{4,5,6,7,8}

		var arr1 = [...]int{1,2,3}
		arr2 := [...]int{4,5,6,7,8}

 Accessing Element of an Array:
	- You can access a specific array element by referring to the index number.
	- Ex.
		prices := [3]int{10,20,30}

		fmt.Println(prices[0]) // Output: 10
		fmt.Println(prices[2]) // Output: 30
 Changing Elements of an Array:
	- You can also change the value of a specific array element by referring to the index number.
	- Ex.
		prices := [3]int{10,20,30}

		prices[2] = 50
		fmt.Println(prices) // Output:[ 10, 20, 50]
 Array Initialization:
	- If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type.
	- The default value for int is 0, and the default value for string is "".
	- Ex.
		arr1 := [5]int{} //not initialized
		arr2 := [5]int{1,2} //partially initialized
		arr3 := [5]int{1,2,3,4,5} //fully initialized

		fmt.Println(arr1) // Output:[ 0 0 0 0 0]
		fmt.Println(arr2) // Output:[ 1 2 0 0 0]
		fmt.Println(arr3) // Output:[ 1 2 3 4 5]
 Initialize Only Specific Elements:
	- It is possible to initialize only specific elements in an array
	- Ex.
		arr1 := [5]int{1:10,2:40}

		fmt.Println(arr1) // Output: [0, 10, 40, 0, 0]
 Finding the length of an Array;
	- The len() function is used to find the length of an array
	- fmt.Println(len(array_name))

 */
// Slices
 /*
 - Slices are similar to arrays, but are more powerful and flexible.
 - used to store multiple values of the same type in a single variable just like arrays
 - But unlike arrays, the length of a slice can grow and shrink as you see fit
 - len() function - returns the length of the slice (the number of elements in the slice)
 - cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

	Creating a Slice with []datatype{values}
		Syntax:
				slice_name := []datatype{}

	Creating a Slice from an Array:
		Syntax:
				var myarray = [length]datatype{values} // An array
				myslice := myarray[start:end] // A slice made from the array

	Creating a Slice with the make() function:
		- this is the correct way to create an empty slice

		Syntax:
				slice_name := make([]type, length, capacity)

	Accessing Slices:
		- You can access a specific slice element by referring to the index number.
		- Ex. 
					prices := []int{10,20,30}

					fmt.Println(prices[0]) // Output: 10
					fmt.Println(prices[2]) // Output: 30

	Changing Slices:
		- You can also change a specific slice element by referring to the index number.
		- Ex.
				prices := []int{10,20,30}
				prices[2] = 50
				fmt.Println(prices[0]) //Output: 10
				fmt.Println(prices[2]) // Output: 50

	Appending Slices:
		- You can append elements to the end of a slice using the append()function:
				Syntax: slice_name = append(slice_name, element1, element2, ...)
		- Ex.
			myslice1 := []int{1, 2, 3, 4, 5, 6}
			fmt.Printf(myslice1) // Output: [1 2 3 4 5 6]
			myslice1 = append(myslice1, 20, 21)
			fmt.Printf(myslice1) // Output:[1 2 3 4 5 6 20 21]

	Appending One Slice to Another Slice:
		- To append all the elements of one slice to another slice, use the append()function

		Syntax:
			slice3 = append(slice1, slice2...)

	Changing the Length of a Slice:
		- It is possible to change the length of a slice.
		- Ex.
			arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
			myslice1 := arr1[1:5] // Slice array
			myslice1 = arr1[1:3] // Change length by re-slicing the array
			myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items

	Memory Efficiency (Copying):
		- When using slices, Go loads all the underlying elements into the memory.
		- If the array is large and you need only a few elements, it is better to copy those elements using the copy() function.
			Syntax:
				copy(dest, src)
		- Ex.
			numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15} // original slice
			neededNumbers := numbers[:len(numbers)-10] // Copy with only need numbers
			copy(numbersCopy, neededNumbers) // Output: [1 2 3 4 5]

 */
// Operators:
 /*
	- The + operator adds together two values
		• var a = 15 + 25
		• sum1 = 100 + 50, sum2 = sum1 + 250
	
	Arithmetic Operators:
		Operator	Name			Description								Example
		+			Addition		Adds together two values				x + y	
		-			Subtraction		Subtracts one value from another		x - y	
		*			Multiplication	Multiplies two values					x * y	
		/			Division		Divides one value by another			x / y	
		%			Modulus			Returns the division remainder			x % y	
		++			Increment		Increases the value of a variable by 1	x++	
		--			Decrement		Decreases the value of a variable by 1	x--
	Assignment Operators( Assignment operators are used to assign values to variables)
		Operator		Example		Same As		
		=				x = 5		x = 5	
		+=				x += 3		x = x + 3	
		-=				x -= 3		x = x - 3	
		*=				x *= 3		x = x * 3	
		/=				x /= 3		x = x / 3	
		%=				x %= 3		x = x % 3	
		&=				x &= 3		x = x & 3	
		|=				x |= 3		x = x | 3	
		^=				x ^= 3		x = x ^ 3	
		>>=				x >>= 3		x = x >> 3	
		<<=				x <<= 3		x = x << 3
	Comparison Operators (Comparison operators are used to compare two values)
		-  The return value of a comparison is either true (1) or false (0).
		Operator			Name						Example	
			==				Equal to					x == y	
			!=				Not equal					x != y	
			>				Greater than				x > y	
			<				Less than					x < y	
			>=				Greater than or equal to	x >= y	
			<=				Less than or equal to		x <= y
	Logical Operators (Logical operators are used to determine the logic between variables or values)
		Operator					Name							Description								Example	
				&& 				Logical and		Returns true if both statements are true					x < 5 &&  x < 10	
				|| 				Logical or		Returns true if one of the statements is true				x < 5 || x < 4	
				!				Logical not		Reverse the result, returns false if the result is true		!(x < 5 && x < 10)
	Bitwise Operators ( Bitwise operators are used on (binary) numbers)
		Operator			Name						Description																Example
			& 				AND							Sets each bit to 1 if both bits are 1									x & y	
			|				OR							Sets each bit to 1 if one of two bits is 1								x | y	
			^				XOR							Sets each bit to 1 if only one of two bits is 1							x ^ b	
			<<				Zero fill left shift		Shift left by pushing zeros in from the right							x << 2	
			>>				Signed right shift			Shift right by pushing copies of the leftmost bit in from the left, 	x >> 2
														and let the rightmost bits fall off	
	
 */
// Conditions
 /*
	- A condition can be either true or false.
	- Go supports the usual comparison operators from mathematics
	- Also supports the usual logical operators
	- Use if to specify a block of code to be executed, if a specified condition is true
	- Use else to specify a block of code to be executed, if the same condition is false
	- Use else if to specify a new condition to test, if the first condition is false
	- Use switch to specify many alternative blocks of code to be executed

	Syntax:
		If statement:
			if condition {
				// code to be executed if condition is true
			}
		If else statement:
			if condition {
			// code to be executed if condition is true
			} else {
			// code to be executed if condition is false
			}
		Else statement:
		 else {
		  // code to be executed if condition is false
		  }
		Else If statement:
			if condition1 {
			// code to be executed if condition1 is true
			} else if condition2 {
			// code to be executed if condition1 is false and condition2 is true
			} else {
			// code to be executed if condition1 and condition2 are both false
			}
		Nested If statement:
			if condition1 {
			// code to be executed if condition1 is true
			if condition2 {
				// code to be executed if both condition1 and condition2 are true
			}
			}
		Switch statement:
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
		Multi-case switch statement:
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
// Loops
 /*
	- The for loop loops through a block of code a specified number of times.
	- The for loop is the only loop available in Go.
	- Loops are handy if you want to run the same code over and over again, each time with a different value.
	- Each execution of a loop is called an iteration

	For Loop:
		Syntax:
			for statement1; statement2; statement3 {
			// code to be executed for each iteration
			}
					1. statement1 Initializes the loop counter value. i := 0
					2. statement2 Evaluated for each loop iteration. If it evaluates to TRUE, the loop continues. If it evaluates to FALSE, the loop ends. i < 5
					3. statement3 Increases the loop counter value. i++
	Continue statement:
		 - The continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.
		 - Ex.
			for i:=0; i < 5; i++ {
				if i == 3 {
				continue
				}
				fmt.Println(i)
				} // Output: 0 1 2 4
	Break statement: 
		- The break statement is used to break/terminate the loop execution.
		- Ex.
			for i:=0; i < 5; i++ {
				if i == 3 {
				break
				}
				fmt.Println(i)	
				}	// Output: 0 1 2
	Nested Loops:
		- It is possible to place a loop inside another loop.
		- Here, the "inner loop" will be executed one time for each iteration of the "outer loop"
		- Ex.
			adj := [2]string{"big", "tasty"}
			fruits := [3]string{"apple", "orange", "banana"}
			for i:=0; i < len(adj); i++ {
				for j:=0; j < len(fruits); j++ {
				fmt.Println(adj[i],fruits[j])
				}
			}
	The Range Keyword:
		- The range keyword is used to more easily iterate through the elements of an array, slice or map. 
		- It returns both the index and the value.
		- Syntax:
				for index, value := range array|slice|map {
				// code to be executed for each iteration
				}
	Omitting Index or Values:
		Index:
				fruits := [3]string{"apple", "orange", "banana"}
				for _, val := range fruits {
					fmt.Printf("%v\n", val)
				}
		Value:
			fruits := [3]string{"apple", "orange", "banana"}
			for idx, _ := range fruits {
				fmt.Printf("%v\n", idx)
			}

 */
// Functions
 /*
 - A function is a block of statements that can be used repeatedly in a program.
 - A function will not execute automatically when a page loads.
 - A function will be executed by a call to the function.

 Creating a Function:
		- Use the func keyword.
		- then followed by (parameters){code}
		- Syntax:
				func FunctionName(para1 type, etc.){
				// code to be executed
				}
		- Call a function:
				• call in the func main(){}
				• call by using functionName()
 Return Values:
	- If you want the function to return a value, you need to define the data type of the return value (such as int, string, etc), and also use the return keyword inside the function
	- Syntax:
				func FunctionName(param1 type, param2 type) type {
				// code to be executed
				return output
				}
	- Named Return Values:
				func myFunction(x int, y int) (result int) {
				result = x + y
				return
				}

				func main() {
				fmt.Println(myFunction(1, 2)) // Output: 3
				}
	- Store the Return Value in a Variable:
				func myFunction(x int, y int) (result int) {
				result = x + y
				return
				}

				func main() {
				total := myFunction(2, 4) // Output: 6
				fmt.Println(total)
				}
	- Multiple Return Values
		func myFunction(x int, y string) (result int, txt1 string) {
		result = x + x
		txt1 = y + " World!"
		return
		}

		func main() {
		fmt.Println(myFunction(5, "Hello")) // Output: 10 Hello World!
		}
	- Omitting Return Values
		- If we (for some reason) do not want to use some of the returned values, we can add an underscore (_), to omit this value.
			func myFunction(x int, y string) (result int, txt1 string) {
				result = x + x
				txt1 = y + " World!"
				return
				}

				func main() {
				_, b := myFunction(5, "Hello")
				fmt.Println(b)
				}
 Recursion:
		- Go accepts recursion functions. 
		- A function is recursive if it calls itself and reaches a stop condition.
				func testcount(x int) int {
				if x == 11 {
					return 0
				}
				fmt.Println(x)
				return testcount(x + 1)
				}

				func main(){
				testcount(1)
				}
	
 */
// Structs
 /*
	- A struct is used to create a collection of members of different data types, into a single variable.
	- While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.

	Declaring a Struct:
		- To declare a structure in Go, use the type and struct keywords:
		- Syntax:
				type struct_name struct {
					member1 datatype;
					member2 datatype;
					member3 datatype;
					...
					}

				then you can use
				var <variableName> Struct
				<variableName>.member = value
 */
// Maps
 /*
	- Maps are used to store data values in key:value pairs.
	- Each element in a map is a key:value pair
	- A map is an unordered and changeable collection that does not allow duplicates.
	- The length of a map is the number of its elements. You can find it using the len() function.
	- The default value of a map is nil.
	- Maps hold references to an underlying hash table.

	Creating Maps using var and := :
		Syntax:
				var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
				b := map[KeyType]ValueType{key1:value1, key2:value2,...}
	
	Creating Maps using make() function:
		Syntax:
					var a = make(map[KeyType]ValueType)
					b := make(map[KeyType]ValueType)

					// Initialize a value
					map_name[KeyType] = value
	Creating a Empty Map:
		Syntax (same as creating a make() map):
				var a = make(map[KeyType]ValueType)
	Allowed Key Types:
		- Everything but slices, maps, and functions
	Access Map Elements:
		Syntax:
			value = map_name[key]
		Ex. 
			var a = make(map[string]string)
				a["brand"] = "Ford"
				a["model"] = "Mustang"
				a["year"] = "1964"

				fmt.Printf(a["brand"]) // Output: Ford
	Updating and Adding Map Elements:
		Syntax:
				map_name[key] = value
	Deleting Map Elements:
		Syntax:
				delete(map_name, key)
	Checking For Specific Elements in a Map
		Syntax:
				val, ok :=map_name[key]
				- If you only want to check the existence of a certain key, you can use the blank identifier (_) in place of val.
	Maps can be References:
		- If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.
		- Ex.
				var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
				b := a

				fmt.Println(a) // Output: [brand:Ford model:Mustang year:1964]
				fmt.Println(b) // Output: [brand:Ford model:Mustang year:1964]

				b["year"] = "1970"
				fmt.Println("After change to b:")

				fmt.Println(a) // Output: [brand:Ford model:Mustang year:1970]
				fmt.Println(b) // Output: [brand:Ford model:Mustang year:1970]
	Iterating Over Maps
		- You can use range to iterate over maps
		- where k = keyType and v = valueType
		- Ex.
				a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

				for k, v := range a {
					fmt.Printf("%v : %v, ", k, v) // Output: two : 2, three : 3, four : 4, one : 1,
				}
	Iterating Over Maps in a Specific Order
		- Maps are unordered data structures.
		- If you need to iterate over a map in a specific order, you must have a separate data structure that specifies that order.
		- Ex. 
				a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

				var b []string             // defining the order
				b = append(b, "one", "two", "three", "four")

				for k, v := range a {        // loop with no order
					fmt.Printf("%v : %v, ", k, v) // Output: two : 2, three : 3, four : 4, one : 1,
				}

				fmt.Println()

				for _, element := range b {  // loop with the defined order
					fmt.Printf("%v : %v, ", element, a[element]) // Output: one : 1, two : 2, three : 3, four : 4,
				}
	 */
// End 