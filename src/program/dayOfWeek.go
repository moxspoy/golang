package main 
import "fmt"

func main() {
	x := 9
	what := ""
	switch (x) {
		case 1,2,3,4,5 : what = "weekday"
		case 6,7 : what = "weekend"
		default : what = "not defined"
	}
	fmt.Println(what)
}