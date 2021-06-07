package golog

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/fatih/color"
)

func (lm *msgLog) control() {
	// format = printFileline() + format // printfileline()打印出错误的文件和行数
	// 判断是输出控制台 还是写入文件
	if lm.out {
		// 如果是输出到控制台，直接执行就好了
		lm.printLine()
		return
	} else {
		// 写入文件
		if everyDay {
			// 如果每天备份的话， 文件名需要更新
			thisDay := fmt.Sprintf("%d-%d-%d", lm.create.Year(), lm.create.Month(), lm.create.Day())
			if lm.now == "" {
				lm.now = thisDay
			}
			if thisDay != lm.now {
				// 重命名
				if err := os.Rename(lm.logPath, filepath.Join(lm.path, lm.now+"_"+lm.name)); err != nil {
					fmt.Println(err)
					lm.out = true
					return
				}
				lm.now = thisDay
			}

		}
		if fileSize > 0 {
			f, err := os.Open(lm.logPath)
			if err == nil {
				// 如果大于设定值， 那么
				fi, err := f.Stat()
				if err == nil && fi.Size() >= fileSize*1024 {
					os.Rename(lm.name, fmt.Sprintf("%d_%s", lm.create.Unix(), lm.name))
				}
				defer f.Close()
			}

		}
		// 如果按照文件大小判断的话，名字不变
		lm.writeToFile()

	}
}

func (lm *msgLog) writeToFile() {
	//
	//if _, ok := logName[name]; !ok {
	//不存在就新建
	f, err := os.OpenFile(lm.logPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		// 如果失败，切换到控制台输出
		color.Red("Permission denied,  auto change to Stdout")
		lm.out = true
		lm.printLine()
		return
	}
	now := lm.create.Format("2006-01-02 15:04:05")
	logMsg := fmt.Sprintf("%s - [%s] - %s - %s - %s - %v\n", now, lm.level, lm.prev, hostname, lm.line, lm.msg)
	f.Write([]byte(logMsg))
	f.Close()
}

func (lm *msgLog) printLine() {
	now := lm.create.Format("2006-01-02 15:04:05")
	color.New(lm.color...).Printf("%s - [%s] - %s - %s - %s - %v\n", now, lm.level, lm.prev, hostname, lm.line, lm.msg)
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
