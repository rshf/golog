package golog

import (
	"testing"
)

func TestNewLog(t *testing.T) {
	defer Sync()
	Info(1111)
	InitLogger("log\\test.log", 0, false)
	Info(2222)
	logging1 := NewLog("log\\test_0.log", 10<<10, false)
	logging2 := NewLog("log\\test_0.log", 10<<10, false)

	logging1.Info("44444444")
	logging2.Info("5555")
}
