package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i) 
  fmt.Print(j)
  // Output: HelloWorld

  fmt.Print(i, "\n")
  fmt.Print(j, "\n")
  /* Output: 
  Hello
  World
  */

  fmt.Print(i, "\n",j)
  /* Output: 
  Hello
  World
  */

  fmt.Print(i, " ", j)
  // Output: Hello World

  fmt.Println(i,j)
  // Output: Hello World

  var h string = "Hello"
  var m int = 15

  // v = value and t = type
  fmt.Printf("h has value: %v and type: %T\n", h, h)
  fmt.Printf("m has value: %v and type: %T", m, m)
  /*
  i has value: Hello and type: string
  j has value: 15 and type: int
  */
} 