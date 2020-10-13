package golog

import (
	"time"
)

var cache chan msgLog

var Synchronous bool

type msgLog struct {
	msg    string
	level  level
	create time.Time
	deep   int
}

func SetSync(synchronous bool) {
	if Synchronous && cache != nil {
		cache = make(chan msgLog, 10000)
		go write()
	}
}

func write() {
	for {
		select {
		case c := <-cache:
			control(c.level, c.msg, c.create, c.deep)
		}
	}
}

// func Sync() {
// 	// 等待日志写完
// 	close(cache)
// }
