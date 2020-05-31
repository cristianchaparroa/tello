package hand

import (
	"fmt"
	"math"

	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
)

// LeftListener is in charge to handle the offets to the left side
type LeftListener struct {
	next EventListener
	c    *tello.Driver
}

// NewLeftListener generates a pointer to the LeftListener
func NewLeftListener(c *tello.Driver) EventListener {
	return &LeftListener{c: c}
}

// Process is in charge to process the left movement
func (l *LeftListener) Process(hand leap.Hand) {

	if l.isLeftEvent(hand) {
		fmt.Println("--> IsLeftMovement")
		l.moveLeft(hand)
		return
	}

	if l.next == nil {
		return
	}

	l.next.Process(hand)

}

// isLeftEvent determines if the current event is for move the drone to the
// left side
func (l *LeftListener) isLeftEvent(hand leap.Hand) bool {
	xAxis := hand.PalmNormal[0]

	// It verifies if there is a shift in the X axis and if it is up
	// to the defined treshold.
	isThreshold := math.Abs(xAxis) > DirectionThreshold

	isUpToZero := xAxis > 0

	// It verifies if the x movement starts in the +X axis.
	return isThreshold && isUpToZero
}

// moveLeft is in charge to calculate the shift and apply it.
func (l *LeftListener) moveLeft(hand leap.Hand) {
	xAxis := hand.PalmNormal[0]
	value := math.Abs(xAxis*10-DirectionThreshold) * DirectionSpeedFactor
	l.c.Left(int(value))
	fmt.Printf(" --> Moving left:%v", value)
}

// SetNext asign the next in the chain of listeners
func (l *LeftListener) SetNext(next EventListener) {
	l.next = next
}
