package readwrite

import (
	"fmt"
	"io"
	"os"
)

// 以流的方式读取文件
func Read(){
	// 获取程序执行的路径
	url,_:=os.Getwd()
	// 打开需要读取的文件
	file,err:=os.Open(url+"/main.go")
	if err!=nil{
		fmt.Println(err)
		return
	}
	// 关闭读取流
	defer file.Close()
	var tmp = make([]byte, 128)
	// 开始读取
	n,err1:=file.Read(tmp)

	//判断是否读取完毕   io.EOF读取完毕
	if err1 ==io.EOF{
		fmt.Println("读取完毕")
		return
	}
	fmt.Println(string(tmp[:n]))
}
