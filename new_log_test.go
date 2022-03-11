package golog

import "testing"

func TestNewLog(t *testing.T) {
	defer Sync()
	logging := NewLog("log\\test_0.log", 10<<10, false)

	logging.Info("44444444")
}
