// Go Access, Change, Append and Copy Slices

/*
Access Elements of a Slice:

package main
import ("fmt")

func main() {
  prices := []int{10,20,30}

  fmt.Println(prices[0])
  fmt.Println(prices[2])
}
// Output: 10 30

*/

/*
Change Elements of a Slice:

package main
import ("fmt")

func main() {
  prices := []int{10,20,30}
  prices[2] = 50
  fmt.Println(prices[0])
  fmt.Println(prices[2])
}
// Output: 10 50

*/

/*
Append Elements To a Slice:

// You can append elements to the end of a slice using the append()function
// Syntax: slice_name = append(slice_name, element1, element2, ...)

package main
import ("fmt")

func main() {
  myslice1 := []int{1, 2, 3, 4, 5, 6}
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  myslice1 = append(myslice1, 20, 21)
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))
}

// Output:
// myslice1 = [1 2 3 4 5 6]
// length = 6
// capacity = 6
// myslice1 = [1 2 3 4 5 6 20 21]
// length = 8
// capacity = 12
*/

/*
Append One Slice To Another Slice

// Syntax: slice3 = append(slice1, slice2...)

package main
import ("fmt")

func main() {
  myslice1 := []int{1,2,3}
  myslice2 := []int{4,5,6}
  myslice3 := append(myslice1, myslice2...)

  fmt.Printf("myslice3=%v\n", myslice3)
  fmt.Printf("length=%d\n", len(myslice3))
  fmt.Printf("capacity=%d\n", cap(myslice3))
}

// Output: 
// myslice3=[1 2 3 4 5 6]
// length=6
// capacity=6

*/

/*
Change The Length of a Slice: Unlike arrays, it is possible to change the length of a slice

package main
import ("fmt")

func main() {
  arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
  myslice1 := arr1[1:5] // Slice array
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  myslice1 = arr1[1:3] // Change length by re-slicing the array
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))
}

// Output: 
// myslice1 = [10 11 12 13]
// length = 4
// capacity = 5
// myslice1 = [10 11]
// length = 2
// capacity = 5
// myslice1 = [10 11 20 21 22 23]
// length = 6
// capacity = 10

*/

/*
Memory Efficiency:

// If the array is large and you need only a few elements
// it is better to copy those elements using the copy() function.

// copy() creates a new underlying array with only the required elements for the slice. 
// This will reduce the memory used for the program. 

// Syntax: copy(dest, src)
// function takes in two slices dest and src, and copies data from src to dest. 
// It returns the number of elements copied.

*/

package main
import ("fmt")

func main() {
  numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
  // Original slice
  fmt.Printf("numbers = %v\n", numbers)
  fmt.Printf("length = %d\n", len(numbers))
  fmt.Printf("capacity = %d\n", cap(numbers))

  // Create copy with only needed numbers
  neededNumbers := numbers[:len(numbers)-10]
  numbersCopy := make([]int, len(neededNumbers))
  copy(numbersCopy, neededNumbers)

  fmt.Printf("numbersCopy = %v\n", numbersCopy)
  fmt.Printf("length = %d\n", len(numbersCopy))
  fmt.Printf("capacity = %d\n", cap(numbersCopy))
}

/*
Output: 
// Original slice
numbers = [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]
length = 15
capacity = 15
// New slice
numbersCopy = [1 2 3 4 5]
length = 5
capacity = 5
*/