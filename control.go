package golog

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/fatih/color"
)

var day = ""

type Format struct {
	lv   level
	msg  interface{}
	deep int
}

func control(lv level, msg interface{}) {
	// format = printFileline() + format // printfileline()打印出错误的文件和行数
	// 判断是输出控制台 还是写入文件
	if stdOut {
		printLine(lv, msg)
		return
	} else {
		// 写入文件
		localtime := time.Now()
		if everyDay {
			// 如果每天备份的话， 文件名需要更新

			thisDay := fmt.Sprintf("%d-%d-%d", localtime.Year(), localtime.Month(), localtime.Day())
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
					os.Rename(Name, fmt.Sprintf("%d_%s", localtime.Unix(), Name))
				}
				defer f.Close()
			}

		}
		// 如果按照文件大小判断的话，名字不变
		thisName := filepath.Join(logPath, Name)
		writeToFile(thisName, lv, msg)
	}
}

func writeToFile(name string, lv level, msg interface{}) {
	//
	//if _, ok := logName[name]; !ok {
	//不存在就新建
	f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		// 如果失败，切换到控制台输出
		color.Red("Permission denied,  auto change to Stdout")
		stdOut = true
		printLine(lv, msg)
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	if lv == Up {
		msg = fmt.Sprintf("caller from %s, msg: %v", printFileline(int(Up)), msg)
	}
	logMsg := fmt.Sprintf("%s - [%s] - %s - %s - %v\n", now, lv, hostname, printFileline(0), msg)

	f.Write([]byte(logMsg))
	f.Close()
}

func writeToFilef(name string, lv level, format string, args ...interface{}) {
	writeToFile(name, lv, fmt.Sprintf(format, args...))
}

func printLine(lv level, msg interface{}) {
	now := time.Now().Format("2006-01-02 15:04:05")
	if lv == Up {
		msg = fmt.Sprintf("caller from %s -- %v", printFileline(int(Up)), msg)
		lv = DEBUG
	}
	color.New(logColor[lv]...).Printf("%s - [%s] - %s - %s - %v\n", now, lv, hostname, printFileline(0), msg)
}

func printLinef(lv level, format string, args ...interface{}) {
	printLine(lv, fmt.Sprintf(format, args...))
}

func printFileline(c int) string {
	c += 4
	_, file, line, ok := runtime.Caller(c)
	if !ok {
		file = "???"
		line = 0
	}
	return fmt.Sprintf("%s:%d", file, line)
}
