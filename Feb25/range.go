/*
The Range Keyword: 
- The range keyword is used to more easily iterate through the elements of an array, slice or map. 
It returns both the index and the value.

Syntax:

for index, value := range array|slice|map {
   // code to be executed for each iteration
}

*/

package main
import ("fmt")

func main() {
  fruits := [3]string{"apple", "orange", "banana"}
  for idx, val := range fruits {
     fmt.Printf("%v\t%v\n", idx, val)
  }

  for _, val := range fruits {
	fmt.Printf("%v\n", val)
 }

 for idx, _ := range fruits {
	fmt.Printf("%v\n", idx)
 }

}

/*
Both Values and Index
Output:
0      apple
1      orange
2      banana

Values:
Output:
apple
orange
banana

Index
Output:
0
1
2
*/
