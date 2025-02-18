// Create a Slice With []datatype{values}
	/* Syntax:
	slice_name := []datatype{values}
				or
	myslice := []int{} // myslice := []int{1,2,3} to initialize/declare
	*/

// Two Helpful Functions in Go
/*
1. len() function - returns the length of the slice 
				(the number of elements in the slice)
2. cap() function - returns the capacity of the slice 
	(the number of elements the slice can grow or shrink to)
*/

package main
import ("fmt")

func main() {
  myslice1 := []int{}
  fmt.Println(len(myslice1))
  fmt.Println(cap(myslice1))
  fmt.Println(myslice1)

  myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
  fmt.Println(len(myslice2))
  fmt.Println(cap(myslice2))
  fmt.Println(myslice2)
}

/* Output:
0
0
[]
4
4
[Go Slices Are Powerful]
*/


