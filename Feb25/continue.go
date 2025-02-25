/*
The Continue Statement:
	- The continue statement is used to skip one or more iterations in the loop. 
	- It then continues with the next iteration in the loop.
*/

// Example

package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    if i == 3 {
      continue
    }
	// 3 will not be skipped 
   fmt.Println(i)
  }
}

/*
Output:
0
1
2
4
*/