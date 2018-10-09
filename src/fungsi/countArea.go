package main 
import (
	"fmt"
	"math"
)

/*
Program ini menerapkan fungsi di Go untuk menghitung luas persegi dan lain lain
 */

func areaPersegi(sisi float64) float64 {
	return sisi*sisi
}

func areaSegitiga(alas float64, tinggi float64) float64 {
	var luas float64 = (alas*tinggi)/2
	return luas
}

func areaCircle (radius float64) float64 {
	var luas float64 = math.Pi * radius * radius
	return luas
}

func main () {
	var area float64 = areaPersegi(4.3)
	var area2 float64 = areaSegitiga(6.0,7.2)
	var area3 float64 = areaPersegi(4.0)
	fmt.Printf("Luas persegi adalah %.1f, dan luas segitiga adalah %.2f, dan luas lingkaran adalah %.2f", area, area2, area3)
}
