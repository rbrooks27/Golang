/*
The continue Statement:
- The continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.
The break Statement:
- The break statement is used to break/terminate the loop execution.
*/

package main

import "fmt"

func main() {
	cont()
	bre()
}

func cont() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}

func bre() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
}

/*
Cont
Output:
0
1
2
4

Bre
Output:
0
1
2
*/


