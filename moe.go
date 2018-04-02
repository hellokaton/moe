package moe

import (
	"fmt"
	"strconv"
	"sync/atomic"
	"time"
)

// ClearLine go to the beggining of the line and clear it
const ClearLine = "\r\033[K"

// returns a valid color's foreground text color attribute
var colorMap = map[string]int{
	"red":     31,
	"green":   32,
	"yellow":  33,
	"blue":    34,
	"magenta": 35,
	"cyan":    36,
	"white":   37,
}

// Moe spinner type
type Moe struct {
	speed  time.Duration
	color  int
	frames []string
	pos    int
	active uint64
	text   string
}

// New Spinner with args
func New(text string) *Moe {
	m := &Moe{
		text: ClearLine + "%s " + text,
	}
	return m.Spinner("dots")
}

// Spinner default spinner
func (m *Moe) Spinner(name string) *Moe {
	tmp := SpinnerMap[name]
	m.speed = tmp.speed
	m.frames = tmp.frames
	return m
}

// Frame custom frames
func (m *Moe) Frame(frames []string) *Moe {
	m.frames = frames
	return m
}

// Color spinner color
func (m *Moe) Color(c string) *Moe {
	m.color = colorMap[c]
	return m
}

// Speed speed
func (m *Moe) Speed(speed time.Duration) *Moe {
	m.speed = speed
	return m
}

// Text set spinner text
func (m *Moe) Text(text string) *Moe {
	m.text = ClearLine + "%s " + text
	return m
}

// Start shows the spinner.
func (m *Moe) Start() *Moe {
	if atomic.LoadUint64(&m.active) > 0 {
		return m
	}
	atomic.StoreUint64(&m.active, 1)
	go func() {
		for atomic.LoadUint64(&m.active) > 0 {
			next := m.next()
			fmt.Printf(m.text, next)
			time.Sleep(m.speed)
		}
	}()
	return m
}

// Stop hide spinner
func (m *Moe) Stop() bool {
	if x := atomic.SwapUint64(&m.active, 0); x > 0 {
		fmt.Printf(ClearLine)
		return true
	}
	return false
}

// next spinner token
func (m *Moe) next() string {
	r := m.frames[m.pos%len(m.frames)]
	m.pos++
	if m.color > 0 {
		return "\033[" + strconv.Itoa(m.color) + "m" + string(r) + "\033[m"
	}
	return string(r)
}
