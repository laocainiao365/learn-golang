package Reflect

import (
	"fmt"
	"reflect"
)

type Student struct {
	UserName string `json:"user_name"`
	Age int `json:"age"`
	Sex string `json:"sex"`
}

func (s Student) Show(){
	fmt.Println("student-show")
}

func (s Student) GetInfo() string {
	var str = fmt.Sprintf("姓名:%v 年龄:%v", s.UserName, s.Age)
	return str
}



func StructReflectType (s interface{}) {
	t := reflect.TypeOf(s)
	//通过Field索引值的方式获取一个属性
	fieldInt := t.Field(0)
	fmt.Println(fieldInt)
	//通过FieldNaem方式获取一个属性
	field,bl := t.FieldByName("UserName")
	if !bl {
		fmt.Println(bl,"-------")
	}
	fmt.Println(field)
	//	获取字段的类型
	fmt.Println(field.Type)
	//	获取字段的名字
	fmt.Println(field.Name)
	// 	获取字段的tag名字
	fmt.Println(field.Tag.Get("json"))

	//	获取结构体中方法 Method(index)
	fmt.Println(t.Method(1))

	//fmt.Println(t.NumMethod())
	//fmt.Println(t.Method(1).Name)
	//
	//fmt.Println(t.Method(1).Type)


}

func Fn(){
	s := Student{
		UserName: "小李",
		Age: 24,
		Sex: "男",
	}
	StructReflectType(s)
}



