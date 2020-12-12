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
	"github.com/hyahm/golog"
)

func main() {
	// 第一个参数是设置日志目录 ， 如果为空，默认显示再控制台
	 // 第二个参数是设置日志切割的大小，0 表示不按照大小切割， 默认单位M，
	 //  第三个事是否每天切割，
	 // 第四个是删除多少天以前的日志， 根据设置的name 来匹配， 0表示不删除
	golog.InitLogger("/data/log", 0, false, 1)  // 如果不写初始化， 默认输出到控制台
	x := 2
	golog.Info(x)
	golog.Infof("%d", x)
	golog.Level = golog.INFO // 设打印的日志级别
	golog.UpFunc()   // 打印调用的方法位置
	golog.Sql()   // 写入文件的时候， 2行输出， 第一行 -- 开头， 方便生成的sql直接可以执行
	golog.Sync()  // 日志改成了异步写入。 想保证日志记录完整, 请在代理退出前执行等待日志写入完成
}
```


