package Reflect

import (
	"fmt"
	"reflect"
)

type Customer struct {
	name string
}

func toValue(c interface{}){
	d := reflect.ValueOf(c)
	// 判断指针类型用Elem()
	if reflect.ValueOf(c).Elem().Kind() == reflect.Int{
		d.Elem().SetInt(399)
	}
	//v := reflect.ValueOf(c)
	//stype := v.Kind()

}


//
func ReflectValue(){
	var n = 100
	toValue(n)
	toValue(&n)//传入指针类型
	fmt.Println(n)
}



func Reflect(){
	var a int = 100
	fmt.Println(reflect.TypeOf(a))

	var s string = "hello golang"
	fmt.Println(reflect.TypeOf(s))

	var b bool = false
	fmt.Println(reflect.TypeOf(b))

	// 使用Kind()方法判断引用类型的底层类型
	var arr = [3]int{12,34,53}
	v := reflect.TypeOf(arr)
	fmt.Println(v) //[3]int

	k := v.Kind()
	fmt.Println(k)//array

	var customer = Customer{
		name: "customer",
	}
	cus := reflect.TypeOf(customer)
	fmt.Println(cus)//Reflect.Customer

	// 获取模块的名称
	cname := cus.Name()
	fmt.Println(cname)//Customer

	// 获取原始类型
	stype := cus.Kind()
	fmt.Println(stype) //struct
}







