package main

import "fmt"

func main() {
	sum := 0
	//Three-component loop
	for i := 0; i < 10; i++ {
		fmt.Println(sum)
		sum += 1
	}
	//Three-component loop lite
	for range 10 {
		fmt.Println(sum)
		sum += 1
	}
	fmt.Println(sum)
	//While loop
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n)
	//Infinity loop
	sumNum := 0
	for {
		sumNum++
		fmt.Println(sumNum)
	}
	//For-each range loop
	strings := []string{"hello", "world"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
	//Exit a loop
	sumLoop := 0
	for i := 0; i < 5; i++ {
		if i%2 != 0 {
			continue
		}
		sumLoop += i
	}
	fmt.Println(sumLoop)
}
