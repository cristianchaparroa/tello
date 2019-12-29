package leap

import "gobot.io/x/gobot/platforms/leap"

// ILeapMotionListener handle the events trigered by leap motion device
type ILeapMotionListener interface {
	// ProcessGestures handle the gestures trigered by leap motion device
	ProcessGestures(g leap.Gesture)
}

type leapMotionListener struct {
}

func (l leapMotionListener) ProcessGestures(g leap.Gesture) {

}
