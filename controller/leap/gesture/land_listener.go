package gesture

import (
	"fmt"
	"tello/controller/leap/core"

	log "github.com/sirupsen/logrus"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
)

// LandListener is in charge to handle the events related
// to land events
type LandListener struct {
	next EventListener
	c    *tello.Driver
}

// NewLandListener generates a pointer to LandListener
func NewLandListener(c *tello.Driver) EventListener {
	return &LandListener{c: c}
}

// Process verifies if is land event and trigger it.
func (l *LandListener) Process(gesture leap.Gesture) bool {

	if l.IsLandEvent(gesture) {
		fmt.Println("IsLand")
		l.c.Land()
		return true
	}

	if l.next == nil {
		return false
	}

	return l.next.Process(gesture)
}

// SetNext saves the next listener in the chain
func (l *LandListener) SetNext(next EventListener) {
	l.next = next
}

// IsLandEvent determines if should to land the drone.
func (l *LandListener) IsLandEvent(gesture leap.Gesture) bool {

	isCircleGesture := core.IsCircleGesture(gesture)
	isCounterClockWise := !core.IsClockWise(gesture)
	isTwoRounds := core.IsTwoRounds(gesture)

	log.WithFields(log.Fields{
		"is_circle":        isCircleGesture,
		"is_counter_clock": isCounterClockWise,
		"is_two_rounds":    isTwoRounds,
	}).Info("is_land_event")

	return isCircleGesture && isCounterClockWise && isTwoRounds
}
