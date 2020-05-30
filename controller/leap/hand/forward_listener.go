package hand

import (
	"fmt"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
	"math"
)

type ForwardListener struct {
	next EventListener
	c    *tello.Driver
}

// NewForwardListener generates a pointer to ForwardListener
func NewForwardListener(c *tello.Driver) *ForwardListener {
	return &ForwardListener{c: c}
}

// Process verifies if is forward event and trigger it.
func (l *ForwardListener) Process(hand leap.Hand) {

	if l.isForwardEvent(hand) {
		l.moveForward(hand)
		return
	}

	if l.next == nil {
		return
	}

	l.next.Process(hand)
}

// isForwardEvent determines if should move to forward the drone.
func (l *ForwardListener) isForwardEvent(hand leap.Hand) bool {
	zAxis := hand.PalmNormal[2]
	isThreshold := math.Abs(zAxis) > DirectionThreshold
	isUpToZero := zAxis > 0
	return isThreshold && isUpToZero
}

func (l *ForwardListener) moveForward(hand leap.Hand) {
	zAxis := hand.PalmNormal[2]
	fmt.Println("-> Move forward")
	value := math.Round(zAxis*10-DirectionThreshold) * DirectionSpeedFactor
	l.c.Forward(int(value))
}

// SetNext saves the next listener in the chain
func (l *ForwardListener) SetNext(next EventListener) {
	l.next = next
}
