/*
Operators are used to perform operations on variables and values.

- The + operator adds together two values it can also be used to add together 
a variable and a value, or a variable and another variable

*/



package main
import ("fmt")

func main() {
  var a = 15 + 25
  fmt.Println(a) // Output: 40

 var(
	sum1 = 100 + 50 // 150 (100 + 50)
    sum2 = sum1 + 250 // 400 (150 + 250)
    sum3 = sum2 + sum2 // 800 (400 + 400)
 )
 fmt.Println(sum2)
 fmt.Println(sum3)
 
}
