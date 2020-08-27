package gesture

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"tello/controller/leap/core"

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
func NewTakeOffListener(c *tello.Driver) EventListener {
	return &TakeOffListener{c: c}
}

// Process verifies if is take off event and trigger it.
func (l *TakeOffListener) Process(gesture leap.Gesture) bool {

	if l.IsTakeOffEvent(gesture) {
		fmt.Println("IsTakeOff")
		l.c.TakeOff()
		return true
	}

	if l.next == nil {
		return false
	}

	return l.next.Process(gesture)
}

// SetNext saves the next listener in the chain
func (l *TakeOffListener) SetNext(next EventListener) {
	l.next = next
}

// IsTakeOffEvent determines if should take off the drone.
func (l *TakeOffListener) IsTakeOffEvent(gesture leap.Gesture) bool {

	isCircleGesture := core.IsCircleGesture(gesture)
	isClockWise := core.IsClockWise(gesture)
	isTwoRounds := core.IsTwoRounds(gesture)

	log.WithFields(log.Fields{
		"is_circle":     isCircleGesture,
		"is_clock_wise": isClockWise,
		"is_two_rounds": isTwoRounds,
	}).Info("is_take_off_event")

	return isCircleGesture && isClockWise && isTwoRounds
}
