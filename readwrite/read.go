package readwrite

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// 以流的方式读取文件
func Read() {
	// 获取程序执行的路径
	url, _ := os.Getwd()
	// 打开需要读取的文件
	file, err := os.Open(url + "/readwrite/index.md")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 关闭读取流
	defer file.Close()

	var str []byte
	var tmp = make([]byte, 128)
	for {
		// 开始读取
		_, err1 := file.Read(tmp)

		//判断是否读取完毕   io.EOF读取完毕
		if err1 == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err1 != nil {
			fmt.Println("读取失败")
			return
		}
		str = append(str, tmp...)
	}

	fmt.Println(string(str))
}

// bufio读取
func Bufio() {
	url, _ := os.Getwd()
	file, err := os.Open(url + "/readwrite/index.md")
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	defer file.Close()

	// 创建一个读取对象
	reader := bufio.NewReader(file)
	//声明一个接受整个文件的切片[]byte
	var count []byte
	// 声明一个每次读取的接受的切片 []byte，并且制定每次读入的大小，这里的大小指的是字节1024 = 1k
	var arrByte = make([]byte, 256)
	for {
		// 以Read方式读取
		n, err1 := reader.Read(arrByte)

		// 以读取字符串方式读取
		//str,err1:=reader.ReadString('\n')
		if err1 == io.EOF { //判断是否读取完毕
			fmt.Println("读取完毕！")
			break
		}
		if err1 != nil {
			fmt.Println("读取失败！")
			return
		}

		count = append(count, arrByte[:n]...)
	}
	//fmt.Println(string(arrByte))
	fmt.Println(string(count))
}

// ioutil方式读取文件
func Ioutil(url string) {
	// 获取路径
	// url,_:=os.Getwd()

	// 读取文件，ioutil.ReadFile() 返回[]byte和err
	file, err := ioutil.ReadFile(url)
	// file,err:=ioutil.ReadFile(url+"/readwrite/index.md")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(file))
}

// CustomerWriteFile以字符串方式写入数据
func CustomerWriteFile() {
	url, _ := os.Getwd()
	url = url + "/readwrite/test.go"
	file, err := os.OpenFile(url, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	defer file.Close()
	for i := 0; i < 10; i++ {
		file.WriteString("package readwrite\r\n")
	}
}

// bufio写入方式
func WriteBufio(){
	url,_ := os.Getwd()
	url = url + "/readwrite/1.txt"

	// 打开要写入的文件，如果不存在就创建os.O_CREATE
	// os.OpenFile("写入文件的路径",写入的模式, 写入的权限)
	// 写入的模式：
	//os.O_CREATE 创建不存在的文件
	//os.O_RDWR 写入的文件可读可写
	//os.O_APPEND 追加
	//os.O_TRUNC  清空
	// os.O_RDONLY 只读
	//os.O_WRONLY 只写
	file,err := os.OpenFile(url, os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)
	if err != nil{
		//fmt.Println("打开文件失败！", err)
		log.Println(err)
		return
	}
	defer file.Close()
	log.Println("error")
	write := bufio.NewWriter(file)
	n,err1 := write.WriteString("你好golang")
	if err1 != nil{
		fmt.Println(err1)
		return
	}
	write.Flush()
	fmt.Println(n)
}









