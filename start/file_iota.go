package main

import "fmt"

func main() {
	//const (
	//	Sunday    = 0
	//	Monday    = 1
	//	Tuesday   = 2
	//	Wednesday = 3
	//	Thursday  = 4
	//	Friday    = 5
	//	Saturday  = 6
	//)
	/*
		iota идентификатор Go используется в объявлениях констант для упрощения определений увеличивающихся чисел.
		Сделаем дни недели с использованием iota - теперь это выглядит проще (особенно если много данных):
	*/

	//const (
	//	Sunday = iota
	//	Monday
	//	Tuesday
	//	Wednesday
	//	Thursday
	//	Friday
	//	Saturday
	//	_  // пропускаем 7
	//	Add
	//)

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	const (
		u         = iota * 42 // u == 0 (индекс - 0, поэтому 0 * 42 = 0)
		v float64 = iota * 42 // v == 42.0 (индекс - 1, поэтому 1.0 * 42 = 42.0)
		w         = iota * 42 // w == 84  (индекс - 2, поэтому 2 * 42 = 84)
	)

	fmt.Println(Sunday)   // вывод 0
	fmt.Println(Saturday) // вывод 6
}
