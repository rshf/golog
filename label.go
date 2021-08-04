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
	LogPath  string
	Create   time.Time
	Label    map[string]string
	Deep     int
	Color    []color.Attribute
	Mu       *sync.RWMutex
	Line     string
	Out      bool
	Path     string
	Dir      string
	Size     int64
	EveryDay bool
	Name     string
	Expire   time.Duration
}

func (l *Label) clean() {
	if name == "" || cleanTime <= 0 {
		return
	}
	for {
		time.Sleep(cleanTime)
		fs, _ := ioutil.ReadDir(l.Dir)
		for _, f := range fs {
			if strings.Contains(f.Name(), l.Name) {
				os.Remove(filepath.Join(logPath, f.Name()))
			}
		}

	}
}

func NewLog(path string, size int64, everyday bool, ct ...time.Duration) *Label {
	var expire time.Duration
	path = filepath.Clean(path)
	if len(ct) > 0 {
		expire = ct[0]
	}
	l := &Label{
		Label:    make(map[string]string),
		Mu:       &sync.RWMutex{},
		Path:     path,
		Size:     size,
		EveryDay: everyday,
		Expire:   expire,
	}
	l.LogPath = path
	l.Dir = filepath.Dir(path)
	l.Name = filepath.Base(path)
	if l.Name != "." {
		go l.clean()
	}
	return l
}

func (l *Label) AddLabel(key, value string) *Label {
	l.Mu.Lock()
	defer l.Mu.Unlock()
	l.Label[key] = value
	return l
}

func (l *Label) DelLabel(key string) *Label {
	l.Mu.RLock()
	defer l.Mu.RUnlock()
	delete(l.Label, key)
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
func (l *Label) Tracef(format string, msg ...interface{}) {
	// Access,
	l.Trace(fmt.Sprintf(format, msg...))
}

// open file，  所有日志默认前面加了时间，
func (l *Label) Debug(msg ...interface{}) {
	// debug,
	if Level <= DEBUG {
		l.s(DEBUG, arrToString(msg...))
	}
}

// open file，  所有日志默认前面加了时间，
func (l *Label) Debugf(format string, msg ...interface{}) {
	// Access,
	l.Debug(fmt.Sprintf(format, msg...))
}

// open file，  所有日志默认前面加了时间，
func (l *Label) Info(msg ...interface{}) {
	if Level <= INFO {
		l.s(INFO, arrToString(msg...))
	}
}
func (l *Label) Infof(format string, msg ...interface{}) {
	// Access,
	l.Info(fmt.Sprintf(format, msg...))
}

// 可以根据下面格式一样，在format 后加上更详细的输出值
func (l *Label) Warn(msg ...interface{}) {
	// error日志，添加了错误函数，
	if Level <= WARN {
		l.s(WARN, arrToString(msg...))
	}
}

func (l *Label) Warnf(format string, msg ...interface{}) {
	// Access,
	l.Warn(fmt.Sprintf(format, msg...))
}

// 可以根据下面格式一样，在format 后加上更详细的输出值
func (l *Label) Error(msg ...interface{}) {
	// error日志，添加了错误函数，
	if Level <= ERROR {
		l.s(ERROR, arrToString(msg...))
	}
}

func (l *Label) Errorf(format string, msg ...interface{}) {
	// Access,
	l.Error(fmt.Sprintf(format, msg...))
}

func (l *Label) Fatal(msg ...interface{}) {
	// error日志，添加了错误函数，
	if Level <= FATAL {
		l.s(FATAL, arrToString(msg...))
	}
	Sync()
	os.Exit(1)
}

func (l *Label) Fatalf(format string, msg ...interface{}) {
	// Access,
	l.Fatal(fmt.Sprintf(format, msg...))
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
	for k, v := range l.Label {
		pre += fmt.Sprintf("[%s = %s]", k, v)
	}
	cache <- msgLog{
		prev:    pre,
		msg:     msg,
		level:   level,
		create:  time.Now(),
		color:   GetColor(level),
		line:    printFileline(0),
		out:     l.Name == "." || l.Name == "",
		path:    l.Dir,
		logPath: l.LogPath,
		name:    l.Name,
		size:    l.Size,
	}
}
