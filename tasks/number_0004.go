package main

import "fmt"

func main() {
	/*
		Дано трехзначное число. Переверните его, а затем выведите.
	*/

	var a int64
	fmt.Scan(&a)
	fmt.Print(a % 10)
	fmt.Print(a / 10 % 10)
	fmt.Print(a / 100)
}
