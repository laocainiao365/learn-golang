# golang中的反射

reflect.TypeOf()判断类型(仅限于简单的数据类型)
```go
var a int = 100
reflect.TypeOf(a) //Int

var s string = "hello golang"
reflect.TypeOf(s) //string

var flag bool = true
reflect.TypeOf(flag) //bool
```
对于引用类型的判断需要用Kind()方法
```go
var arr = [3]int{12,34,53}
v := reflect.TypeOf(arr)
// 这种情况下无法准确判断 arr的类型，需要用Kind()
fmt.Println(reflect.TypeOf(v) //[3]int

// 使用Kind()方法，这个可以获取底层原始类型
k := v.Kind()
fmt.Println(k) //array
```
Name()方法可以获取，模块的具体名字
```go
type Customer struct {
	name string
}
...
...
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
```
reflect.ValueOf(params)返回一个reflect.value这个里面有 int()\string()等方法可以将interface{}类型转化为响应的类型进行使用。
```go
func toValue(c interface{}){
    d := reflect.ValueOf(c)
	// 判断指针类型用Elem()
	if reflect.ValueOf(c).Elem().Kind() == reflect.Int{
		d.Elem().SetInt(399)
	}
    v := reflect.ValueOf(c)//这里v是个interface{}类型  不能直接使用 
	stype := v.Kind() // 获取参数的原始类型
	fmt.Println(stype) //int
	sum := v.Int() + 20 // 将空接口类型转化int
	fmt.Println(sum) // 120
}

func ReflectValue(){
	var n = 100
	toValue(n) 
	toValue(&n)//传入指针类型
	fmt.Println(n)
}
```


