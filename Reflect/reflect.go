package Reflect

import (
	"fmt"
	"reflect"
)
func Reflect(){
	var a int = 100
	fmt.Println(reflect.TypeOf(a))

	var s string = "hello golang"
	fmt.Println(reflect.TypeOf(s))

	var b bool = false
	fmt.Println(reflect.TypeOf(b))

	var arr [3]int
	arr[0] = 2
	arr[1]= 34
	arr[2] = 321
	fmt.Println(reflect.TypeOf(arr).Kind())


}







