/*
The break Statement:
	- The break statement is used to break/terminate the loop execution.
*/

// Example
package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    if i == 3 {
      break
    }
	// the program will now exit/terminate because it hits 3
   fmt.Println(i)
  }
}

/*
Output:
0
1
2
*/