# time

time的使用
```golang
//获取时间
time.Now() //2020-03-21 23:38:27.87641 +0800 CST m=+0.000057573

//获取格式化日期
time.Now().Format() //2020/03/21 11:45:13

//获取时间戳
time.Now().Unix()//1592754392

//把时间戳 格式化为时间
var f int64 = 1592754392
s:=time.Unix(f, 0).Format("2006-01-02 03:04:05") //2020-03-21 11:46:32

//把字符串转换为时间对象
var str = "2011-09-03 12:45:42"
var tem = "2006-01-02 03:04:05"
obj, _ := time.ParseInLocation(tem, str, time.Local)
```
