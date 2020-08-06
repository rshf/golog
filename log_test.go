package golog

import (
	"fmt"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	fmt.Println(time.Now())
	Info("two")
	Error("one")
}
