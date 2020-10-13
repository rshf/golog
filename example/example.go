package main

import (
	"github.com/hyahm/golog"
)

type level int

const (
	TRACE level = iota * 10
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	SQL
)

var Level level

func (l level) String() string {
	switch l {
	case 0:
		return "TRACE"
	case 10:
		return "DEBUG"
	case 20:
		return "INFO"
	case 30:
		return "WARN"
	case 40:
		return "ERROR"
	case 50:
		return "FATAL"
	case 60:
		return "SQL"
	default:
		return "DEBUG"
	}
}

func main() {
	golog.InitLogger("log", 0, false)
	golog.Name = "bbbb.log"
	golog.SetSync(true)
	golog.Infof("adf%s", "cander")
	golog.Debug("debug wo和 ")
	golog.Level = golog.TRACE
	golog.Error("wo和 ")
	golog.Trace("7777")
	golog.Warn("warnning")
	golog.SetSync(true)
	golog.Info("999")

	golog.UpFunc(1, "msg")
	golog.Info("me")
}
