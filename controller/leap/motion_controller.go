package leap

import (
	"fmt"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
	"tello/controller/leap/frame"
	"tello/controller/leap/gesture"
	"tello/controller/leap/hand"
	"tello/controller/leap/utils"
)

const (
	// DefaultShift is the shift value in case that there is not a configuration.
	DefaultShift = 20
)

// MotionController is a MotionController in charge to handle the Tello drone movements
type MotionController struct {
	drone  *tello.Driver
	leap   *leap.Driver
	logger utils.Logger
}

// NewController generates pointer to MotionController
func NewController(drone *tello.Driver, leap *leap.Driver) *MotionController {
	logger := utils.NewLogger()
	c := &MotionController{drone: drone, leap: leap, logger: logger}
	return c
}

// Run function create the infinity loop to read all Frames until the user
// wants to land the drone
func (c *MotionController) Run() {

	c.leap.On(leap.GestureEvent, func(data interface{}) {
		g := data.(leap.Gesture)
		c.logger.ShowGesture(g)
		manager := gesture.NewManager(c.drone)
		manager.ProcessGestures(g)
	})

	c.leap.On(leap.HandEvent, func(data interface{}) {
		h := data.(leap.Hand)
		c.logger.ShowHand(h)
		manager := hand.NewManager(c.drone)
		manager.ProcessHands(h)
	})

	c.leap.On(leap.MessageEvent, func(data interface{}) {
		f := data.(leap.Frame)
		openHand :=  frame.IsOpenHand(f.Pointables)
		fmt.Printf("Is open hand:%v \n", openHand)
		//c.logger.ShowFingers(f)
	})
}