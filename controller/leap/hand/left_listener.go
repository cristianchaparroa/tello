package hand

import (
	log "github.com/sirupsen/logrus"
	"math"
	"tello/controller/leap/core"

	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
)

// LeftListener is in charge to handle the offets to the left side
type LeftListener struct {
	next   EventListener
	c      *tello.Driver
	logger core.Logger
}

// NewLeftListener generates a pointer to the LeftListener
func NewLeftListener(c *tello.Driver) EventListener {
	logger := core.NewLogger()
	return &LeftListener{c: c, logger: logger}
}

// Process is in charge to process the left movement
func (l *LeftListener) Process(hand leap.Hand) {

	if l.isLeftEvent(hand) {
		l.logger.ShowHand(hand, "--> IsLeftMovement")
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
	// to the defined threshold.
	isThreshold := math.Abs(xAxis) > DirectionThreshold

	isUpToZero := xAxis > 0

	// It verifies if the x movement starts in the +X axis.
	return isThreshold && isUpToZero
}

// moveLeft is in charge to calculate the shift and apply it.
func (l *LeftListener) moveLeft(hand leap.Hand) {

	xNormal := hand.PalmNormal[0]
	offset := l.calculateShift(xNormal)
	l.c.Left(offset)

	log.WithFields(log.Fields{
		"offset": offset,
	}).Info("moving_left")

}

func (l *LeftListener) calculateShift(shift float64) int {
	offset := math.Abs(shift*50+DirectionThreshold) * DirectionSpeedFactor
	return int(offset)
}

// SetNext asign the next in the chain of listeners
func (l *LeftListener) SetNext(next EventListener) {
	l.next = next
}
