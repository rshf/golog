package golog

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/fatih/color"
)

var day = ""

// type format struct {
// 	lv   level
// 	msg  interface{}
// 	deep int
// }

func (lm *msgLog) control() {
	// format = printFileline() + format // printfileline()打印出错误的文件和行数
	// 判断是输出控制台 还是写入文件

	if stdOut {
		lm.printLine()
		return
	} else {
		// 写入文件
		if everyDay {
			// 如果每天备份的话， 文件名需要更新

			thisDay := fmt.Sprintf("%d-%d-%d", lm.create.Year(), lm.create.Month(), lm.create.Day())
			if day == "" {
				day = thisDay
			}
			if thisDay != day {
				// 重命名
				os.Rename(Name, day+"_"+Name)
				day = thisDay
			}

		} else if fileSize > 0 {

			f, err := os.Open(filepath.Join(logPath, Name))
			if err == nil {
				// 如果大于设定值， 那么
				fi, err := f.Stat()
				if err == nil && fi.Size() >= fileSize*1024*1024 {
					os.Rename(Name, fmt.Sprintf("%d_%s", lm.create.Unix(), Name))
				}
				defer f.Close()
			}

		}
		// 如果按照文件大小判断的话，名字不变
		thisName := filepath.Join(logPath, Name)
		lm.writeToFile(thisName)
	}
}

func (lm *msgLog) writeToFile(name string) {
	//
	//if _, ok := logName[name]; !ok {
	//不存在就新建
	f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		// 如果失败，切换到控制台输出
		color.Red("Permission denied,  auto change to Stdout")
		stdOut = true
		lm.printLine()
		return
	}
	now := lm.create.Format("2006-01-02 15:04:05")
	// if len(deep) > 0 {
	// 	msg = fmt.Sprintf("caller from %s, msg: %v", printFileline(deep[0]), msg)
	// }
	logMsg := fmt.Sprintf("%s - [%s] - %s - %s - %v\n", now, lm.level, hostname, lm.line, lm.msg)
	// cache <- msgLog{
	// 	f:   f,
	// 	msg: logMsg,
	// }
	f.Write([]byte(logMsg))
	f.Close()
}

func (lm *msgLog) printLine() {
	now := lm.create.Format("2006-01-02 15:04:05")

	color.New(lm.color...).Printf("%s - [%s] - %s - %s - %v\n", now, lm.level, hostname, lm.line, lm.msg)
}

func (lm *msgLog) printLinef(lv level, format string, args ...interface{}) {
	lm.printLine()
}

func printFileline(c int) string {
	c += 3
	_, file, line, ok := runtime.Caller(c)
	if !ok {
		file = "???"
		line = 0
	}
	return fmt.Sprintf("%s:%d", file, line)
}
