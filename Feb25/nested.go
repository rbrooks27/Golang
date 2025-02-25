/*
Nested Loops:
- It is possible to place a loop inside another loop.
- Here, the "inner loop" will be executed one time for each iteration of the "outer loop"
*/

package main
import ("fmt")

func main() {
  adj := [2]string{"big", "tasty"}
  fruits := [3]string{"apple", "orange", "banana"}
  for i:=0; i < len(adj); i++ {
    for j:=0; j < len(fruits); j++ {
      fmt.Println(adj[i],fruits[j])
    }
  }
}

/*
Output:
big apple
big orange
big banana
tasty apple
tasty orange
tasty banana
*/