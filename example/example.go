package main

import (
	"github.com/fatih/color"
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
	golog.InitLogger("log", 0, true)
	golog.Level = golog.DEBUG
	c := make([]color.Attribute, 0)
	c = append(c, color.FgBlue)
	golog.SetColor(golog.DEBUG, c)
	golog.Infof("adf%s", "cander")
	golog.Debug("debug wo和 ")
	golog.Level = golog.TRACE
	golog.Error("wo和 ")
	golog.Trace("7777")
	golog.Warn("warnning")
	golog.Info("999")
	aaa()
	golog.Info("me")
	golog.Sync()
}

func aaa() {
	golog.UpFunc(1, "msg")
}
