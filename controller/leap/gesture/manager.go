package gesture

import (
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
)

// Manager handle the gesture triggered by leap motion device
type Manager interface {
	// Process handle the gesture triggered by leap motion device
	Process(g leap.Gesture) bool
}

type leapGestureManager struct {
	c     *tello.Driver
	event EventListener
}

// NewManager listen the events related to leap device.
func NewManager(c *tello.Driver) Manager {
	listener := &leapGestureManager{c: c}
	listener.build()
	return listener
}

func (l *leapGestureManager) build() {
	takeOff := NewTakeOffListener(l.c)
	land := NewLandListener(l.c)
	land.SetNext(takeOff)
	l.event = land
}

func (l *leapGestureManager) Process(gesture leap.Gesture) bool {
	return l.event.Process(gesture)
}
