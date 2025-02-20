/*
For Loops:
- The for loop loops through a block of code a specified number of times.
- The for loop is the only loop available in Go.

Syntax:
for statement1; statement2; statement3 {
   // code to be executed for each iteration
}

1. statement1 Initializes the loop counter value.
2. statement2 Evaluated for each loop iteration. If it evaluates to TRUE, the loop continues. If it evaluates to FALSE, the loop ends.
3. statement3 Increases the loop counter value.

*/

package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }

  for i:=0; i <= 100; i+=10 {
    fmt.Println(i)
  }
}

/*
For 1st
Output:
0
1
2
3
4

For 2nd
Output:
0
10
20
30
40
50
60
70
80
90
100
*/

