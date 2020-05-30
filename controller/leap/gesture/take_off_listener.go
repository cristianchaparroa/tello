package gesture

import (
	"fmt"

	"gobot.io/x/gobot/platforms/dji/tello"

	"gobot.io/x/gobot/platforms/leap"
)

// TakeOffListener is in charge to handle the events related
// to take off events
type TakeOffListener struct {
	next EventListener
	c    *tello.Driver
}

// NewTakeOffListener generates a pointer to TakeOffListener
func NewTakeOffListener(c *tello.Driver) *TakeOffListener {
	return &TakeOffListener{c: c}
}

// Process verifies if is take off event and trigger it.
func (l *TakeOffListener) Process(gesture leap.Gesture) {

	if l.IsTakeOffEvent(gesture) {
		fmt.Println("IsTakeOff")
		l.c.TakeOff()
		return
	}

	if l.next == nil {
		return
	}

	l.next.Process(gesture)
}

// SetNext saves the next listener in the chain
func (l *TakeOffListener) SetNext(next EventListener) {
	l.next = next
}

// IsTakeOffEvent determines if should take off the drone.
func (l *TakeOffListener) IsTakeOffEvent(gesture leap.Gesture) bool {
	isCircleGesture := IsCircleGesture(gesture)
	isClockWise := isClockWise(gesture)
	isTwoRounds := isTwoRounds(gesture)
	fmt.Println(isCircleGesture, isClockWise, isTwoRounds)
	return isCircleGesture && isClockWise
}
