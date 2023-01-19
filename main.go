package main

import (
	"fmt"
	reflect2 "reflect"
)

func main() {
	//check(1)
	//check("1")
	//check(true)
	//check([]int{1, 1, 3, 4})
	fnSlice()

}
func fnSlice() {
	var slice []string
	slice = make([]string, 0)
	slice = append(slice, "1", "2", "3", "4")
	for i := range slice {
		fmt.Println(slice[i])
	}
}
func check(unknown interface{}) {
	ref := reflect2.ValueOf(unknown)
	switch ref.Kind() {
	case reflect2.String:
		var v string = ref.String()
		println(v)
		println("type:string value:", ref.String())
	case reflect2.Int:
		println("type:int value:", ref.Int())
	case reflect2.Bool:
		println("type:bool value:", ref.Bool())
	case reflect2.Slice:
		println("is slice")
		fmt.Println(ref.Slice(0, ref.Len()))
	}
}
