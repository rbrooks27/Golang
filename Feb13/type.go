package main

import "fmt"

func main() {
    var a bool = true     // Boolean
    var b int = 5         // Integer
    var c float32 = 3.14  // Floating point number
    var d string = "Hi!"  // String

    fmt.Printf("Boolean: %v, Type: %T\n", a, a)
    fmt.Printf("Integer: %v, Type: %T\n", b, b)
    fmt.Printf("Float:   %v, Type: %T\n", c, c)
    fmt.Printf("String:  %v, Type: %T\n", d, d)
}
