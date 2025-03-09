package main

import "fmt"

func main() {
	numberArray()
	arraySlice()
	maps()
}

func numberArray() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	numbers := [6]int{2, 2, 2, 3, 4, 5}
	fmt.Println(numbers)
}

func arraySlice() {
	array := [6]int{0, 1, 2, 3, 4, 5}
	var s []int = array[0:3]
	fmt.Println(s)
}

func maps() {
	a := make(map[string]string)
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"
	fmt.Println(a)
	b := make(map[string]int)
	b["Oslo"] = 1
	b["Bergen"] = 2
	b["Trondheim"] = 3
	b["Stavanger"] = 4
	fmt.Println()

}
