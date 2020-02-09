package hand

import (
	"fmt"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
	"math"
)

type BackwardListener struct {
	next EventListener
	c    *tello.Driver
}

// NewBackwardListener generates a pointer to BackwardListener
func NewBackwardListener(c *tello.Driver) *BackwardListener {
	return &BackwardListener{c: c}
}

// Process verifies if is backward event and trigger it.
func (l *BackwardListener) Process(hand leap.Hand) {

	if l.isBackwardEvent(hand) {
		l.moveBackward(hand)
		return
	}

	if l.next == nil {
		return
	}

	l.next.Process(hand)
}
func (l *BackwardListener) isBackwardEvent(hand leap.Hand) bool {
	zAxis := hand.PalmNormal[2]
	isThreshold := math.Abs(zAxis) > DirectionThreshold
	isUpToZero := zAxis > 0
	return isThreshold && !isUpToZero
}

func (l *BackwardListener) moveBackward(hand leap.Hand) {
	fmt.Println("--> Move backward")
	zAxis := hand.PalmNormal[2]
	value := math.Round(zAxis * 10 - DirectionThreshold) * DirectionSpeedFactor
	l.c.Backward(int(value))
}

// SetNext saves the next listener in the chain
func (l *BackwardListener) SetNext(next EventListener) {
	l.next = next
}