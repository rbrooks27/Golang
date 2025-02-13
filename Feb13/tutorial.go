package main
import "fmt"

func main() {
	var student1 string = "John" //type is string
  	var student2 = "Jane" //type is inferred
	var A int = 10 
	x := 2 // can only be used inside fuctions


	fmt.Println(student1) // Output: John
	fmt.Println(student2) /* Output: Jane */
	fmt.Println(A) // Output: 10
    fmt.Println(x) /* Output: 2 */
}

