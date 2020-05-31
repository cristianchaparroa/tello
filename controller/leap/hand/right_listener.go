package hand

import (
	"fmt"
	"math"

	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
)

// RightListener is in charge to handle the offets to the right side
type RightListener struct {
	next EventListener
	c    *tello.Driver
}

// NewRightListener generates a pointer to the RightListener
func NewRightListener(c *tello.Driver) EventListener {
	return &RightListener{c: c}
}

// Process is in charge to process the right movement
func (l *RightListener) Process(hand leap.Hand) {

	if l.isRightEvent(hand) {
		fmt.Println("--> IsRightMovement")
		l.moveRight(hand)
		return
	}

	if l.next == nil {
		return
	}

	l.next.Process(hand)
}

// isRightEvent determines if the current event is for move the drone to the
// right side
func (l *RightListener) isRightEvent(hand leap.Hand) bool {
	xHand := hand.PalmNormal[0]

	// It verifies if there is a shift in the X axis and if it is up
	// to the defined treshold.
	isThreshold := math.Abs(xHand) > DirectionThreshold

	// It verifies if the x movement starts in the -X axis.
	isDownToZero := xHand < 0

	return isThreshold && isDownToZero
}

// moveRight is in charge to calculate the shift and apply it.
func (l *RightListener) moveRight(hand leap.Hand) {
	xAxis := hand.PalmNormal[0]
	value := math.Abs(xAxis*50+DirectionThreshold) * DirectionSpeedFactor
	l.c.Left(int(value))
	fmt.Printf(" --> Moving right:%v", value)
}

// SetNext asign the next in the chain of listeners
func (l *RightListener) SetNext(next EventListener) {
	l.next = next
}
