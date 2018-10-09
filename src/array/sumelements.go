package main 
import "fmt"
func main () {
	a := [...] float64 {2.3, 4.3, 99.5, 8.0}
	var sum float64

	//conventional loop
	for i := 0; i<len(a); i++ {
		//print all
		fmt.Printf("%.2f, ", a[i])
		sum += a[i]
	}
	fmt.Printf("sum of all element is %.2f \n", sum)

	//loop using range
	for index, value := range a {
		fmt.Printf("value in index-%d is %.3f \n", index, value)
	}
}