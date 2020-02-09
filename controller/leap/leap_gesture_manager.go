package leap

import (
	"gobot.io/x/gobot/platforms/leap"
)

type leapGestureManager struct {
	c     *Controller
	event EventListener
}

// NewleapGestureManager listen the events related to leap device.
func NewleapGestureManager(c *Controller) GestureManager {
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
