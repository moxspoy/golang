package main 
import "fmt"
/*
This program will print odd numbers
 */
func main () {
	for i := 1; i < 11; i++ {
		if (i%2 == 0) {
			continue
		}
		fmt.Println(i)
	}
}