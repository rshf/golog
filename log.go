package golog

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	logPath  string = ""    // 文件路径
	fileSize int64  = 0     // 切割的文件大小
	everyDay bool   = false // 每天一个来切割文件 （这个比上面个优先级高）
	stdOut   bool   = true
)

// 文件名
var Name = "info.log"

// hostname
var hostname = ""

func init() {
	hostname, _ = os.Hostname()
}

func InitLogger(path string, size int64, everyday bool) {
	if path != "" {
		stdOut = false
		logPath = filepath.Clean(path)
		err := os.MkdirAll(logPath, 0755)
		if err != nil {
			panic(err)
		}
		fileSize = size
		everyDay = everyday
	}

}

// open file，  所有日志默认前面加了时间，
func Tracef(format string, args ...interface{}) {
	Trace(fmt.Sprintf(format, args...))
}

// open file，  所有日志默认前面加了时间，
func Debugf(format string, args ...interface{}) {
	Debug(fmt.Sprintf(format, args...))
}

// open file，  所有日志默认前面加了时间，
func Infof(format string, args ...interface{}) {
	Info(fmt.Sprintf(format, args...))
}

// 可以根据下面格式一样，在format 后加上更详细的输出值
func Warnf(format string, args ...interface{}) {
	// error日志，添加了错误函数，
	Warn(fmt.Sprintf(format, args...))
}

// 可以根据下面格式一样，在format 后加上更详细的输出值
func Errorf(format string, args ...interface{}) {
	// error日志，添加了错误函数，
	Error(fmt.Sprintf(format, args...))
}

func Fatalf(format string, args ...interface{}) {
	// error日志，添加了错误函数，
	Fatal(fmt.Sprintf(format, args...))
}

func UpFuncf(deep int, format string, args ...interface{}) {
	// deep打印函数的深度， 相对于当前位置向外的深度
	UpFunc(deep, fmt.Sprintf(format, args...))
}

// open file，  所有日志默认前面加了时间，
func Trace(msg ...interface{}) {
	// Access,
	if Level <= TRACE {
		control(TRACE, arrToString(msg...), time.Now())
	}
}

// open file，  所有日志默认前面加了时间，
func Debug(msg ...interface{}) {
	// debug,
	if Level <= DEBUG {
		s(Level, arrToString(msg...))
	}
}

// open file，  所有日志默认前面加了时间，
func Info(msg ...interface{}) {
	if Level <= INFO {
		s(Level, arrToString(msg...))
	}
}

// 可以根据下面格式一样，在format 后加上更详细的输出值
func Warn(msg ...interface{}) {
	// error日志，添加了错误函数，
	if Level <= WARN {
		s(Level, arrToString(msg...))
	}
}

// 可以根据下面格式一样，在format 后加上更详细的输出值
func Error(msg ...interface{}) {
	// error日志，添加了错误函数，
	if Level <= ERROR {
		s(Level, arrToString(msg...))
	}
}

func Fatal(msg ...interface{}) {
	// error日志，添加了错误函数，
	if Level <= FATAL {
		s(Level, arrToString(msg...))
	}
	os.Exit(1)
}

func UpFunc(deep int, msg ...interface{}) {
	// deep打印函数的深度， 相对于当前位置向外的深度
	if Level <= DEBUG {
		s(DEBUG, arrToString(msg...))
	}
}

func arrToString(msg ...interface{}) string {
	ll := make([]string, 0, len(msg))
	for range msg {
		ll = append(ll, "%v")
	}
	return fmt.Sprintf(strings.Join(ll, ""), msg...)
}

func s(level level, msg string, deep ...int) {
	if Synchronous {
		cache <- msgLog{
			msg:    msg,
			level:  level,
			create: time.Now(),
		}
	} else {
		control(INFO, msg, time.Now())
	}
}
