# golog
### 安装
```
 go get github.com/hyahm/golog
```

### 使用
> 在main 函数开始直接调用InitLogger方法  
```
package main

import (
	"app"
	"github.com/hyahm/golog"
)

func main() {
	golog.InitLogger("/data/log", 0, false)
	x := 2
	golog.Info(x)
	golog.Infof("%d", x)
	golog.Level = golog.INFO // 设打印的日志级别
	golog.UpFunc()   // 打印调用的方法位置
	golog.Sql()   // 写入文件的时候， 2行输出， 第一行 -- 开头， 方便生成的sql直接可以执行
	golog.Sync()  // 请在代理退出前执行等待日志写入完成
}
```
### 直接指定配置(日志文件夹路径, 分割日志的大小, 单位Mb, 是否每天分割日志)  , 只是控制台输出， 可以不用调用InitLogger()
```
golog.Name = "", 设置日志文件名
golog.InitLogger("/data/log", 0, false)   // 如果路径为空， 默认输出到控制台
```

> 后面任何地方直接调用方法即可写入, 每个方法都会生成一个文件   
```
golog.Info(format string, args ...interface{})
golog.Debug(format string, args ...interface{})
golog.Fatal(format string, args ...interface{})
golog.Error(format string, args ...interface{})
```


