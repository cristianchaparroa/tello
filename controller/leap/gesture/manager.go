package gesture

import (
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
)

// GestureManager handle the gesture trigered by leap motion device
type GestureManager interface {
	// ProcessGestures handle the gesture trigered by leap motion device
	ProcessGestures(g leap.Gesture)
}


type leapGestureManager struct {
	c     *tello.Driver
	event EventListener
}

// NewManager listen the events related to leap device.
func NewManager(c *tello.Driver) GestureManager {
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

func (l *leapGestureManager) ProcessGestures(gesture leap.Gesture) {
	l.event.Process(gesture)
}