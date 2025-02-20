/*
Assignment Operators
- Assignment operators are used to assign values to variables.
*/

package main
import ("fmt")

func main() {

// The assignment operator (=) to assign a value

  var x = 10
  fmt.Println(x)

  var y = 15
  y +=5
  fmt.Println(y)

}

/* 

A list of all assignment operators:

Operator		Example		Same As		
=				x = 5		x = 5	
+=				x += 3		x = x + 3	
-=				x -= 3		x = x - 3	
*=				x *= 3		x = x * 3	
/=				x /= 3		x = x / 3	
%=				x %= 3		x = x % 3	
&=				x &= 3		x = x & 3	
|=				x |= 3		x = x | 3	
^=				x ^= 3		x = x ^ 3	
>>=				x >>= 3		x = x >> 3	
<<=				x <<= 3		x = x << 3

*/