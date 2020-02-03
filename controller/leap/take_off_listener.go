package leap

import (
	"fmt"

	"gobot.io/x/gobot/platforms/leap"
)

// TakeOffListener is in charge to handle the events related
// to take off events
type TakeOffListener struct {
	next EventListener
	c    *Controller
}

// NewTakeOffListener generates a pointer to TakeOffListener
func NewTakeOffListener(c *Controller) *TakeOffListener {
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

// IsTakeOffEvent determins if is  take off event.
func (l *TakeOffListener) IsTakeOffEvent(gesture leap.Gesture) bool {
	isCircleGeture := IsCircleGesture(gesture)
	isClockWise := isClockWise(gesture)
	isTwoRounds := isTwoRounds(gesture)

	return isCircleGeture && isClockWise && isTwoRounds
}
