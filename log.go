package golog

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	logPath  string = ""    // 文件路径
	fileSize int64  = 0     // 切割的文件大小
	everyDay bool   = false // 每天一个来切割文件 （这个比上面个优先级高）
	stdOut   bool   = true
)

// 文件名
var Name = "info.log"
var mu sync.Mutex

var hostname = ""

func init() {
	hostname, _ = os.Hostname()
}

func InitLogger(path string, size int64, everyday bool) {
	mu = sync.Mutex{}
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
	// Access,
	if Level <= TRACE {
		controlf(TRACE, format, args...)
	}
}

// open file，  所有日志默认前面加了时间，
func Debugf(format string, args ...interface{}) {
	// debug,
	if Level <= DEBUG {
		controlf(DEBUG, format, args...)
	}
}

// open file，  所有日志默认前面加了时间，
func Infof(format string, args ...interface{}) {
	// debug,
	if Level <= INFO {
		controlf(INFO, format, args...)
	}
}

// 可以根据下面格式一样，在format 后加上更详细的输出值
func Warnf(format string, args ...interface{}) {
	// error日志，添加了错误函数，
	if Level <= WARN {
		controlf(WARN, format, args...)
	}
}

// 可以根据下面格式一样，在format 后加上更详细的输出值
func Errorf(format string, args ...interface{}) {
	// error日志，添加了错误函数，
	if Level <= ERROR {
		controlf(ERROR, format, args...)
	}
}

func Fatalf(format string, args ...interface{}) {
	// error日志，添加了错误函数，
	if Level <= FATAL {
		controlf(FATAL, format, args...)
	}
}

func Sqlf(format string, args ...interface{}) {
	// error日志，添加了错误函数，
	if Level <= SQL {
		controlf(SQL, format, args...)
	}
}

// open file，  所有日志默认前面加了时间，
func Trace(msg ...interface{}) {
	// Access,
	if Level <= TRACE {
		control(TRACE, arrToString(msg...))
	}
}

// open file，  所有日志默认前面加了时间，
func Debug(msg ...interface{}) {
	// debug,
	if Level <= DEBUG {
		control(DEBUG, arrToString(msg...))
	}
}

// open file，  所有日志默认前面加了时间，
func Info(msg ...interface{}) {
	// debug,
	if Level <= INFO {
		control(INFO, arrToString(msg...))
	}
}

// 可以根据下面格式一样，在format 后加上更详细的输出值
func Warn(msg ...interface{}) {
	// error日志，添加了错误函数，
	if Level <= WARN {
		control(WARN, arrToString(msg...))
	}
}

// 可以根据下面格式一样，在format 后加上更详细的输出值
func Error(msg ...interface{}) {
	// error日志，添加了错误函数，
	if Level <= ERROR {
		control(ERROR, arrToString(msg...))
	}
}

func Fatal(msg ...interface{}) {
	// error日志，添加了错误函数，
	if Level <= FATAL {
		control(FATAL, arrToString(msg...))
	}
}

func Sql(msg ...interface{}) {
	// error日志，添加了错误函数，
	if Level <= SQL {
		control(SQL, arrToString(msg...))
	}
}

func UpFunc(msg interface{}) {
	control(up, msg)
}

func arrToString(msg ...interface{}) string {
	ll := make([]string, 0, len(msg))
	for range msg {
		ll = append(ll, "%v")
	}
	return fmt.Sprintf(strings.Join(ll, ""), msg...)
}
