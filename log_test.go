package golog

import (
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	defer Sync()
	InitLogger("aaaa/gggg.log", 0, false)
	Info(4444)
	t.Log("com")
	time.Sleep(time.Second * 6)
}
