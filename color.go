package golog

import (
	"sync"

	"github.com/fatih/color"
)

var logColor map[level][]color.Attribute
var cmu *sync.RWMutex

func init() {
	cmu = &sync.RWMutex{}
	logColor = make(map[level][]color.Attribute)
	logColor[ERROR] = []color.Attribute{color.FgRed}
}

func SetColor(lv level, attrs []color.Attribute) {
	cmu.Lock()
	logColor[lv] = attrs
	cmu.Unlock()
}

func DelColor(lv level, attrs []color.Attribute) {
	cmu.Lock()
	delete(logColor, lv)
	cmu.Unlock()
}
