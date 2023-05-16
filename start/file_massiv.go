package main

import "fmt"

func main() {
	//var a [3]int
	//fmt.Println(a)

	// ****************************************
	var a [3]int = [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [...]int{1, 2, 3}
	d := [3]int{1: 12}

	fmt.Println(a) // [1 2 3]
	fmt.Println(b) // [1 2 3]
	fmt.Println(c) // [1 2 3]
	fmt.Println(d) // [0 12 0]
	// *****************************************

	a1 := [3]int{1, 2, 3}
	b1 := [3]int{1, 2, 3}
	c1 := [3]int{3, 2, 1}

	fmt.Println(a1 == b1) // true
	fmt.Println(a1 == c1) // false
	// *****************************************

	var numbers [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println(numbers[0]) // 1
	fmt.Println(numbers[4]) // 5

	numbers[0] = 87

	fmt.Println(numbers[0]) // 87
	// *****************************************

	/*
		Range возвращает 2 объекта: индекс элемента в массиве и копию значения этого элемента. Любой из этих объектов должен
		быть опущен, если мы не планируем использовать его, для этого вместо имени переменной мы можем указать символ _.
		Кроме того, если мы хотим использовать только индекс элемента, мы можем вообще не использовать второе возвращаемое
		значение:
	*/

	a2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a) // [1 2 3 4 5]

	for idx, elem := range a2 {
		fmt.Printf("Элемент с индексом %d: %d\n", idx, elem)
		// Элемент с индексом 0: 1
		// Элемент с индексом 1: 2
		// Элемент с индексом 2: 3
		// Элемент с индексом 3: 4
		// Элемент с индексом 4: 5
	}

	a3 := [5]int{1, 2, 3, 4, 5}

	for idx := range a3 {
		fmt.Println(a[idx])
	}

	for idx, _ := range a3 {
		// В этом случае следует использовать приведенный выше вариант,
		// хотя технически эти варианты работают одинаково
		fmt.Println(a[idx])
	}

	for _, elem := range a3 {
		fmt.Println(elem)
	}

	/*
		Необходимо запомнить, что в качестве второго значения range возвращает копию элемента массива, это может быть важно,
		если в цикле мы хотим изменить массив. В этом случае мы должны обращаться к элементам массива по индексу:
	*/

	a4 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a) // [1 2 3 4 5]

	for _, elem := range a4 {
		elem = 100
		fmt.Println(elem)

		// 100
		// 100
		// 100
		// 100
		// 100
	}
	fmt.Println(a4) // [1 2 3 4 5]

	for idx := range a4 {
		a[idx] = 100
		fmt.Println(a[idx])

		// 100
		// 100
		// 100
		// 100
		// 100
	}
	fmt.Println(a4) // [100 100 100 100 100]

}