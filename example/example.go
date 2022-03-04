package main

import (
	"github.com/hyahm/golog"
)

func main() {
	defer golog.Sync()
	logger1 := golog.NewLog("log/test1.log", 0, false)
	logger2 := golog.NewLog("log/test2.log", 0, false)
	logger1.Info("foo")
	logger2.Info("foo")
}
