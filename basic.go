package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-future/model"
	reflect2 "reflect"
)

const (
	index = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func fnConst() {
	println(Monday, Tuesday, Wednesday)
}
func fnFunc() {
	x, y := 1, 2
	println(x, y)
	swap(&x, &y)
	println(&x, &y)
	println(x, y)

}
func swap(x *int, y *int) {
	*x = *x ^ *y
	*y = *x ^ *y
	*x = *x ^ *y
}
func printErrMsg(f func()) (status int, msg string) {
	defer func() {
		//设置recover拦截错误信息
		err := recover()
		//产生panic异常  打印错误信息
		if err != nil {
			status = 400
			msg = "error"
		}
	}()
	f()
	status = 200
	msg = "success"
	return
}
func fnDefer() {
	var arr [10]int
	var len = 10
	x := func() {
		arr[len] = 10
	}
	status, msg := printErrMsg(x)
	println(status, msg)

	println(x)
}

func fnList() {
	var array [3]int32
	array[0] = 1
	array[1] = 2
	array[2] = 3
	var arr = [4]int{1, 2, 3, 4}
	for i := range arr {
		print(arr[i])
	}
}

func fnOOP() {
	//var u = model.User{
	//	Name: "sss",
	//}
	//
	//fmt.Printf("%+v\n", u)
	//fmtPrintln(u)
	//var son = model.NewSon(model.NewParent("王淼", "男"), "walking")

	//println(son.Say())
	//tood(son)
	//tood(model.User{})
	//var p = model.NewParent("李荣", "女")
	//p.Walk()
	//son.Walk()
	//son.Parent.Walk()
	//println(son.Name)
	var _obj model.Action
	_obj = model.NewSon(model.NewParent("王淼", "男"), "walking")
	println(_obj.Run())
	_obj = model.User{}
	println(_obj.Run())
	//println(assets(11))
	println(fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", 31, "vvv"))
	//fmt.Printf("\x1b[%dmhello world 30: 黑 \x1b[0m\n", 30)
	//fmt.Printf("\x1b[%dmhello world 31: 红 \x1b[0m\n", 31)
	//fmt.Printf("\x1b[%dmhello world 32: 绿 \x1b[0m\n", 32)
	//fmt.Printf("\x1b[%dmhello world 33: 黄 \x1b[0m\n", 33)
	//fmt.Printf("\x1b[%dmhello world 34: 蓝 \x1b[0m\n", 34)
	//fmt.Printf("\x1b[%dmhello world 35: 紫 \x1b[0m\n", 35)
	//fmt.Printf("\x1b[%dmhello world 36: 深绿 \x1b[0m\n", 36)
	//fmt.Printf("\x1b[%dmhello world 37: 白色 \x1b[0m\n", 37)
	//
	//fmt.Printf("\x1b[%d;%dmhello world \x1b[0m 47: 白色 30: 黑 \n", 47, 30)
	//fmt.Printf("\x1b[%d;%dmhello world \x1b[0m 46: 深绿 31: 红 \n", 46, 31)
	//fmt.Printf("\x1b[%d;%dmhello world \x1b[0m 45: 紫   32: 绿 \n", 45, 32)
	//fmt.Printf("\x1b[%d;%dmhello world \x1b[0m 44: 蓝   33: 黄 \n", 44, 33)
	//fmt.Printf("\x1b[%d;%dmhello world \x1b[0m 43: 黄   34: 蓝 \n", 43, 34)
	//fmt.Printf("\x1b[%d;%dmhello world \x1b[0m 42: 绿   35: 紫 \n", 42, 35)
	//fmt.Printf("\x1b[%d;%dmhello world \x1b[0m 41: 红   36: 深绿 \n", 41, 36)
	//fmt.Printf("\x1b[%d;%dmhello world \x1b[0m 40: 黑   37: 白色 \n", 40, 37)

}
func tood(action model.Action) {
	println(action.Say())
}

func assets(args interface{}) (bool, string) {
	value, ok := args.(string)
	if !ok {
		var msg = fmt.Sprintf("type error!need string,but %v", 0x4B, reflect2.TypeOf(args), 0x4B)
		panic(msg)
	}
	return ok, value
}

type Son struct {
	name string
}

func fmtPrintln(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(v)
		return
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	if err != nil {
		fmt.Println(v)
		return
	}

	fmt.Println(out.String())
}

func reflect() {
	var num float64 = 1.2345

	fmt.Println("type: ", reflect2.TypeOf(num))
	fmt.Println("value: ", reflect2.ValueOf(num))
}
