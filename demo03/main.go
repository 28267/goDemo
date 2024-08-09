package main

import (
	"fmt"
	"reflect"
)

func test(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.String {
		v1 := v.String()
		fmt.Printf("类型是%T,值=%v\n", v1, v1)
	}
	if v.Kind() == reflect.Int {
		v1 := v.Int()
		fmt.Printf("类型是%T,值=%v\n", v1, v1)
	}

}

func main() {
	var str = "hello"
	var num = 15
	test(str)
	test(num)
}
