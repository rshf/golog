package golog

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/fatih/color"
)

func TestLog(t *testing.T) {
	defer Sync()
	Info("11111")
	SetColor(INFO, []color.Attribute{
		color.BgBlack,
		color.FgGreen,
	})
	Info("22222")
	Info("3333")
	Error("4444")

	Info("5555")
}

func TestLog1(t *testing.T) {
	defer Sync()
	InitLogger("log\\bbbbb.log", 0, true)
	Info("11111")
	SetColor(INFO, []color.Attribute{
		color.BgBlack,
		color.FgGreen,
	})
	Info("22222")
	Info("3333")
	Error("4444")

	Info("5555")
}

func TestColor(t *testing.T) {
	attrs := []color.Attribute{
		color.FgBlue,
		color.Bold,
	}
	logPath = filepath.Clean("")
	color.New(attrs...).Println(logPath)
}

func TestNewLog(t *testing.T) {
	defer Sync()
	a := NewLog("log\\aaa.log", 0, false)
	a.Info("bbb")
	a.Info("aaaa")
	a.Info("cccc")
	a.Info("ddddd")
}

func TestFilePath(t *testing.T) {
	x := filepath.Join("log", "bb.log")
	t.Log(x)
	os.Rename("log\\aaaa.log", x)
	t.Log(filepath.Dir(""))
	t.Log(filepath.Base(""))
}
