package golog

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

type Label struct {
	msg      string
	level    level
	create   time.Time
	label    map[string]string
	deep     int
	color    []color.Attribute
	mu       *sync.RWMutex
	line     string
	out      bool
	path     string
	dir      string
	size     int64
	everyDay bool
	name     string
	expire   time.Duration
}

func (l *Label) clean() {
	if name == "" || cleanTime <= 0 {
		return
	}
	for {
		time.Sleep(cleanTime)
		fs, _ := ioutil.ReadDir(l.dir)
		for _, f := range fs {
			if strings.Contains(f.Name(), l.name) {
				os.Remove(filepath.Join(logPath, f.Name()))
			}
		}

	}
}

func NewLog(path string, size int64, everyday bool, ct ...time.Duration) *Label {
	var expire time.Duration

	if len(ct) > 0 {
		expire = ct[0]
	}
	l := &Label{
		label:    make(map[string]string),
		mu:       &sync.RWMutex{},
		path:     path,
		size:     size,
		everyDay: everyday,
		expire:   expire,
	}
	return l
}

func (l *Label) AddLabel(key, value string) *Label {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.label[key] = value
	return l
}

func (l *Label) DelLabel(key string) *Label {
	l.mu.RLock()
	defer l.mu.RUnlock()
	delete(l.label, key)
	return l
}

// open file，  所有日志默认前面加了时间，
func (l *Label) Trace(msg ...interface{}) {
	// Access,
	if Level <= TRACE {
		l.s(TRACE, arrToString(msg...))
	}
}

// open file，  所有日志默认前面加了时间，
func (l *Label) Debug(msg ...interface{}) {
	// debug,
	if Level <= DEBUG {
		l.s(DEBUG, arrToString(msg...))
	}
}

// open file，  所有日志默认前面加了时间，
func (l *Label) Info(msg ...interface{}) {
	if Level <= INFO {
		l.s(INFO, arrToString(msg...))
	}
}

// 可以根据下面格式一样，在format 后加上更详细的输出值
func (l *Label) Warn(msg ...interface{}) {
	// error日志，添加了错误函数，
	if Level <= WARN {
		l.s(WARN, arrToString(msg...))
	}
}

// 可以根据下面格式一样，在format 后加上更详细的输出值
func (l *Label) Error(msg ...interface{}) {
	// error日志，添加了错误函数，
	if Level <= ERROR {
		l.s(ERROR, arrToString(msg...))
	}
}

func (l *Label) Fatal(msg ...interface{}) {
	// error日志，添加了错误函数，
	if Level <= FATAL {
		l.s(FATAL, arrToString(msg...))
	}
	Sync()
	os.Exit(1)
}

func (l *Label) UpFunc(deep int, msg ...interface{}) {
	// deep打印函数的深度， 相对于当前位置向外的深度
	if Level <= DEBUG {
		l.s(DEBUG, arrToString(msg...), deep)
	}
}

func (l *Label) s(level level, msg string, deep ...int) {
	if len(deep) > 0 && deep[0] > 0 {
		msg = fmt.Sprintf("caller from %s -- %v", printFileline(deep[0]), msg)
	}
	pre := ""
	for k, v := range l.label {
		pre += fmt.Sprintf("[%s = %s]", k, v)
	}
	cache <- msgLog{
		prev:   pre,
		msg:    msg,
		level:  level,
		create: time.Now(),
		color:  GetColor(level),
		line:   printFileline(0),
		out:    logPath == "",
	}
}
