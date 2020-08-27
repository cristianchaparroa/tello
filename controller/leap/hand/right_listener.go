package hand

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"math"
	"tello/controller/leap/core"

	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
)

// RightListener is in charge to handle the offets to the right side
type RightListener struct {
	next   EventListener
	c      *tello.Driver
	logger core.Logger
}

// NewRightListener generates a pointer to the RightListener
func NewRightListener(c *tello.Driver) EventListener {
	logger := core.NewLogger()
	return &RightListener{c: c, logger: logger}
}

// Process is in charge to process the right movement
func (l *RightListener) Process(hand leap.Hand) bool {

	if l.isRightEvent(hand) {
		fmt.Println("--> IsRightMovement")
		l.logger.ShowHand(hand)
		l.moveRight(hand)
		return true
	}

	if l.next == nil {
		return false
	}

	return l.next.Process(hand)
}

// isRightEvent determines if the current event is for move the drone to the
// right side
func (l *RightListener) isRightEvent(hand leap.Hand) bool {
	xNormal := hand.PalmNormal[0]
	isThreshold := IsXMovement(xNormal)
	// It verifies if the x movement starts in the -X axis.
	isDownToZero := xNormal < 0

	return isThreshold && isDownToZero
}

// moveRight is in charge to calculate the shift and apply it.
func (l *RightListener) moveRight(hand leap.Hand) {

	xNormal := hand.PalmNormal[0]
	offset := l.calculateShift(xNormal)
	l.c.Left(offset)

	log.WithFields(log.Fields{
		"offset": offset,
	}).Info("moving_right")
}

func (l *RightListener) calculateShift(shift float64) int {
	offset := math.Abs(shift*50-DirectionThreshold) * DirectionSpeedFactor
	return int(offset)
}

// SetNext assign the next in the chain of listeners
func (l *RightListener) SetNext(next EventListener) {
	l.next = next
}
