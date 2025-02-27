/*
Maps Using make() function:
		- creates a empty map 
		- you can change it from there
	*/

	/*
	Syntax:
		var a = make(map[KeyType]ValueType)
					or
		b := make(map[KeyType]ValueType)
	*/

	/*
	Create an Empty Map:
		var a map[KeyType]ValueType

		Example:
			package main
			import ("fmt")

			func main() {
			var a = make(map[string]string)
			var b map[string]string

			fmt.Println(a == nil)
			fmt.Println(b == nil)
			} // Output: false true
	*/

	// Example
		package main
		import ("fmt")

		func main() {
		var a = make(map[string]string) // The map is empty now
		a["brand"] = "Nike"
		a["model"] = "Paul George 6"
		a["year"] = "2021"
										// a is no longer empty
		b := make(map[string]int)
		b["Stephon Castle"] = 1
		b["Wemby"] = 2
		b["Paolo"] = 3
		b["Lamelo"] = 4

		fmt.Printf("a\t%v\n", a)
		fmt.Printf("b\t%v\n", b)
		}

	/*
	Output:
		a       map[brand:Nike model:Paul George 6 year:2021]
		b       map[Lamelo:4 Paolo:3 Stephon Castle:1 Wemby:2]
	*/

	/*
	Allowed Key Types:
		- Booleans
		- Numbers
		- Strings
		- Arrays
		- Pointers
		- Structs
		- Interfaces (as long as the dynamic type supports equality)
		(The map key can be of any data type for which the equality operator (==) is defined all above)


		Invalid Key Types:
		- Slices
		- Maps
		- Functions
		(These types are invalid because the equality operator (==) is not defined for them.)
	*/

	/*
	Allowed Value Types:
		- The map values can be any type
	*/