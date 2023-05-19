package main

import "fmt"

func main() {
	/*
		Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0. Напишите программу, которая выведет
		элементы массива, индексы которых четны (0, 2, 4...).
	*/

	var a uint64
	fmt.Scan(&a)
	workArray := make([]int, a)
	for i := 0; i < int(a); i++ {
		fmt.Scan(&workArray[i])
	}
	for i := 1; i < int(a+1); i += 2 {
		fmt.Print(workArray[i])
	}
}
