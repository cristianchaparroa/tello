package leap

import (
	"fmt"

	"gobot.io/x/gobot/platforms/leap"
)

// LandListener is in charge to handle the events related
// to land events
type LandListener struct {
	next EventListener
	c    *Controller
}

// NewLandListener generates a pointer to LandListener
func NewLandListener(c *Controller) *LandListener {
	return &LandListener{c: c}
}

// Process verifies if is land event and trigger it.
func (l *LandListener) Process(gesture leap.Gesture) {

	if l.IsLandEvent(gesture) {
		fmt.Println("IsLand")
		l.c.Land()
		return
	}

	if l.next == nil {
		return
	}

	l.next.Process(gesture)
}

// SetNext saves the next listener in the chain
func (l *LandListener) SetNext(next EventListener) {
	l.next = next
}

// IsLandEvent determins if is  land event.
func (l *LandListener) IsLandEvent(gesture leap.Gesture) bool {
	isCircleGeture := IsCircleGesture(gesture)
	isCounterClockWise := !isClockWise(gesture)
	isTwoRounds := isTwoRounds(gesture)

	return isCircleGeture && isCounterClockWise && isTwoRounds
}
