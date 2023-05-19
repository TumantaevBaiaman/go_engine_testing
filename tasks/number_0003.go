package main

import "fmt"

func main() {
	/*
		Дано трехзначное число. Найдите сумму его цифр.
	*/

	var a, sum_number int64
	fmt.Scan(&a)
	for a > 10 {
		sum_number += a % 10
		a = a / 10
	}
	fmt.Println(sum_number)
}
