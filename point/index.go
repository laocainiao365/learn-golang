package point

import "fmt"

func Point(c interface{}) {
	fmt.Println(c)
	c[0] = "sdfdfs"
}

func Show(){
	//var a int = 100
	//Point(&a)
	//fmt.Println(a)

	//var s string = "golang"
	//Point(&s)
	//fmt.Println(s)

	var arr = [2]string{"123","456"}
	Point(&arr)
	fmt.Println(arr)


}
