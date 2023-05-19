package main

import "fmt"

func main() {
	/*
		Дана последовательность, состоящая из целых чисел. Напишите программу, которая подсчитывает количество
		положительных чисел среди элементов последовательности.
	*/

	var a uint64
	var count_num uint64 = 0
	fmt.Scan(&a)
	workArray := make([]int, a)
	for i := 0; i < int(a); i++ {
		fmt.Scan(&workArray[i])
		if workArray[i] >= 0 {
			count_num++
		}
	}
	fmt.Print(count_num)
}
