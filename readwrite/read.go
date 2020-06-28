package readwrite

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 以流的方式读取文件
func Read(){
	// 获取程序执行的路径
	url,_:=os.Getwd()
	// 打开需要读取的文件
	file,err:=os.Open(url+"/readwrite/index.md")
	if err!=nil{
		fmt.Println(err)
		return
	}
	// 关闭读取流
	defer file.Close()

	var str []byte
	var tmp = make([]byte, 128)
	for  {
		// 开始读取
		_,err1:=file.Read(tmp)

		//判断是否读取完毕   io.EOF读取完毕
		if err1 ==io.EOF{
			fmt.Println("读取完毕")
			break
		}
		if err1 != nil{
			fmt.Println("读取失败")
			return
		}
		str = append(str, tmp...)
	}

	fmt.Println(string(str))
}

// bufio读取
func Bufio(){
	url,_:=os.Getwd()
	file,err:=os.Open(url+"/readwrite/index.md")
	if err !=nil{
		fmt.Println("打开文件失败")
		return
	}

	reader:=bufio.NewReader(file)
	var arrByte = make([]byte, 1024)

	n,err1:=reader.Read(arrByte)
	//str,err1:=reader.ReadString('\n')
	if err1 !=nil{
		fmt.Println("读取失败！")
		return
	}

	//fmt.Println(string(arrByte))
	fmt.Println(n)
}


