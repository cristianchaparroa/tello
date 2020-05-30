package hand

import (
	"fmt"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
	"math"
)

type LeftListener struct {
	next EventListener
	c    *tello.Driver
}

func NewLeftListener(c *tello.Driver) *LeftListener {
	return &LeftListener{c: c}
}

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
func (l *LeftListener) isLeftEvent(hand leap.Hand) bool {
	xAxis := hand.PalmNormal[0]
	isThreshold := xAxis > DirectionThreshold
	isUpToZero := xAxis > 0
	return isThreshold && isUpToZero
}

func (l *LeftListener) moveLeft(hand leap.Hand) {
	xAxis := hand.PalmNormal[0]
	value := math.Abs(xAxis*10+DirectionThreshold) * DirectionSpeedFactor
	l.c.Left(int(value))
	fmt.Printf(" --> Moving left:%v", value)
}

func (l *LeftListener) SetNext(next EventListener) {
	l.next = next
}
