package main 
import (
	"fmt"
	"math"
)

func main () {
	var a, b = 4, 5
	var rst = (a+b) * (a/b)
	var r float64 = 7
	var c float64= math.Pi * r * r
	fmt.Printf("a is %v and b is %v, the result is %d, the area of circle is %v", a,b,rst,c)
}