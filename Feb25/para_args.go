/*
Parameters and Arguments:
	- Parameters act as variables inside the function.
	- Parameters and their types are specified after the function name, inside the parentheses. 
	  You can add as many parameters as you want, just separate them with a comma.


Syntax:
	func FunctionName(param1 type, param2 type, param3 type) {
	// code to be executed
	}

*/

// Ex. Function with Parameter and Function with Multiple Parameters
	package main
	import ("fmt")

	func familyName(fname string){
		fmt.Println("Hello", fname, "Refsnes")
	}

	func rivalFamilyName(rname string, age int){
		fmt.Println("Hi", age, "year old", rname, "Drake")
	}

	func main(){
	familyName("Liam") // parameter is fname 
	familyName("Jenny") // parameter is fname 
	familyName("Anja") // parameter is fname 
	rivalFamilyName("Lake", 17)
	rivalFamilyName("Jeffery", 39)
	rivalFamilyName("Marissa", 37)

  /*
	Output:
		Hello Liam Refsnes
		Hello Jenny Refsnes
		Hello Anja Refsnes 
	*/
/*
	Output: 
		Hi 17 year old Lake Drake
		Hi 39 year old Jeffery Drake
		Hi 37 year old Marissa Drake
*/
}


