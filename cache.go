package golog

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
)

var cache chan msgLog

type msgLog struct {
	msg    string
	level  level
	create time.Time
	deep   int
	color  []color.Attribute
	line   string
	out    bool
}

func init() {
	cache = make(chan msgLog, 1000)
	exit = make(chan bool)
	go write()
	go clean()
}

var exit chan bool

func clean() {
	if logPath == "" || cleanTime <= 0 {
		return
	}
	for {
		select {
		case <-time.After(cleanTime):
			fs, _ := ioutil.ReadDir(logPath)
			for _, f := range fs {
				if strings.Contains(f.Name(), Name) {
					os.Remove(filepath.Join(logPath, f.Name()))
				}
			}

		}

	}
}

func write() {
	for c := range cache {
		c.control()
	}
	exit <- true
}

func Sync() {
	// 等待日志写完
	close(cache)
	<-exit
}
