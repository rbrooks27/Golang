/* Ways to declare an array

With the var keyword:

var array_name = [length]datatype{values} // here length is defined

or

var array_name = [...]datatype{values} // here length is inferred

With the := 

array_name := [length]datatype{values} // here length is defined

or

array_name := [...]datatype{values} // here length is inferred

*/

package main
import ("fmt")

func main() {
  var arr1 = [3]int{1,2,3}
  arr2 := [5]int{4,5,6,7,8}

  fmt.Println(arr1)
  fmt.Println(arr2)

  /* 
Output:
[1 2 3]
[4 5 6 7 8]
*/
}

/*
Array of Strings:

package main
import ("fmt")

func main() {
  var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  fmt.Print(cars)
}

//Output: [Volvo BMW Ford Mazda]


*/

/*
Accessing Elements of Array:

package main
import ("fmt")

func main() {
  prices := [3]int{10,20,30}

  fmt.Println(prices[0])
  fmt.Println(prices[2])
}

// Output: 
// 10 
// 30

*/

/*
Changing Elements of an Array:

package main
import ("fmt")

func main() {
  prices := [3]int{10,20,30}

  prices[2] = 50
  fmt.Println(prices)
}

// Output: [10, 20, 50]

*/

/*
Array Initialization:

package main
import ("fmt")

func main() {
  arr1 := [5]int{} //not initialized
  arr2 := [5]int{1,2} //partially initialized
  arr3 := [5]int{1,2,3,4,5} //fully initialized

  fmt.Println(arr1)
  fmt.Println(arr2)
  fmt.Println(arr3)
}

// Output: 
// [0 0 0 0 0]
// [1 2 0 0 0]
// [1 2 3 4 5]

*/

/*
Initialize Only Specific Elements:

package main
import ("fmt")

func main() {
  arr1 := [5]int{1:10,2:40}

  fmt.Println(arr1)
}

// Output: [0 10 40 0 0]

*/

/*
Finding the Length of an Array:

package main
import ("fmt")

func main() {
  arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  arr2 := [...]int{1,2,3,4,5,6}

  fmt.Println(len(arr1))
  fmt.Println(len(arr2))
}

// Output: 4 6
*/


