package leap

import (
	"fmt"

	"gobot.io/x/gobot/platforms/leap"
)

const (
	TurnThreshold        = 0.2
	TurnSpeedFactor      = 2.0
	DirectionThreshold   = 0.25
	DirectionSpeedFactor = 0.05
	UpControlThreshold   = 50
	UpSpeedFactor        = 0.01
	// CircleThreshold is the threshold to determine that the gesture is two rounds
	CircleThreshold = 1.5
)

type leapMotionListener struct {
	c     *Controller
	event EventListener
}

// NewleapMotionListener listen the events related to leap device.
func NewleapMotionListener(c *Controller) MotionListener {
	listener := &leapMotionListener{c: c}
	listener.build()
	return listener
}

func (l *leapMotionListener) build() {
	takeOff := NewTakeOffListener(l.c)
	land := NewLandListener(l.c)
	land.SetNext(takeOff)
	l.event = land
}

func (l *leapMotionListener) ProcessGestures(gesture leap.Gesture) {

	fmt.Println("--> ProcessGestures")
	l.event.Process(gesture)

	fmt.Println("<-- ProcessGestures")
}
