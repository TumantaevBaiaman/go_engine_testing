package main

import "fmt"

func main() {
	var a float64
	fmt.Scan(&a)
	if a > 0 && a <= 10000 {
		fmt.Printf("%.4f", a*a)
	} else if a > 10000 {
		fmt.Printf("%e", a)
	} else {
		fmt.Printf("%s, %2.2f, %s", "число", a*a, "не подходит")
	}
}
