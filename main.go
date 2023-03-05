package main

import (
	"fmt"
)

func same[T int | string, E bool](a T, b T) E {
	return a == b
}
func closure(string2 string) func() int {
	num := 0
	return func() int {
		fmt.Print("入参为:\t" + string2)
		num++
		return num
	}
}
func main() {
	fmt.Println(same("1", "1"))
}
