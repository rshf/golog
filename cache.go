package golog

import (
	"time"

	"github.com/fatih/color"
)

var cache chan msgLog

var Synchronous bool

type msgLog struct {
	msg    string
	level  level
	create time.Time
	deep   int
	color  []color.Attribute
	line   string
}

func init() {
	cache = make(chan msgLog, 1000)
	exit = make(chan bool)
	go write()
}

var exit chan bool

func write() {
	for c := range cache {
		c.control()
	}
	exit <- true
	// for {
	// 	select {
	// 	case c := <-cache:
	// 		fmt.Println("0000")
	// 		c.control()
	// 	}
	// }
}

func Sync() {
	// 等待日志写完
	close(cache)
	<-exit
}
