package main

import (
	"fmt"

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
	golog.Name = "aaa.log"
	cmd := fmt.Sprintf(`ruby ruby/addudid.rb "%s" "%s" "%s" "%s" "%s"`, "fgjhgjg", "dosj@#%ASF1", "Seng's Ipad", "aaaa", "3")
	golog.Info(cmd)
	golog.Infof("adf%s", "cander")
	golog.Debug("wo和 ")
	golog.Level = golog.TRACE
	golog.Error("wo和 ")
	golog.Fatal("asdf")
	golog.Sql("aaa", "bbb", "ddd")
	w()
}

func w() {
	golog.UpFunc()
	golog.Info("me")
}
