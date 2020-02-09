package hand

import (
	"fmt"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
	"math"
)

type RightListener struct {
	next EventListener
	c    *tello.Driver
}

func NewRightListener(c  *tello.Driver) *RightListener {
	return &RightListener{c: c}
}

func (l *RightListener) Process(hand leap.Hand) {

	if l.isRightEvent(hand) {
		fmt.Println("IsRightMovement")
		l.c.Land()
		return
	}

	if l.next == nil {
		return
	}

	l.next.Process(hand)
}

func (l *RightListener) isRightEvent(hand leap.Hand) bool {
	xAxis := hand.PalmNormal[0]
	isThreshold := xAxis > DirectionThreshold
	isUpToZero := xAxis > 0
	return isThreshold && isUpToZero
}

func (l *RightListener) moveRight(hand leap.Hand) {
	xAxis := hand.PalmNormal[0]
	value := math.Abs(xAxis * 10 - DirectionThreshold) * DirectionSpeedFactor
	l.c.Left(int(value))
	fmt.Printf(" --> Moving right:%v", value)
}


func (l *RightListener) SetNext(next EventListener) {
	l.next = next
}
