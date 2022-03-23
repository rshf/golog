package golog

import (
	"os"
	"testing"
)

func TestNewLog(t *testing.T) {
	defer Sync()
	logging := NewLog("log\\test_0.log", 10<<10, false)

	logging.Info("44444444")
}

func TestMkdir(t *testing.T) {
	err := os.MkdirAll("log", 0755)
	t.Log(err)
}
