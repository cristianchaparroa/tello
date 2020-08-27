package leap

import (
	"fmt"
	"tello/controller/leap/core"
	"tello/controller/leap/gesture"
	"tello/controller/leap/hand"

	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
)

const (
	// DefaultShift is the shift value in case that there is not a configuration.
	DefaultShift = 20
)

// MotionController is a MotionController in charge to handle the Tello drone movements
type MotionController struct {
	drone  *tello.Driver
	leap   *leap.Driver
	logger core.Logger
}

// NewLeapMotion generates pointer to MotionController
func NewLeapMotion(drone *tello.Driver, leap *leap.Driver) *MotionController {
	logger := core.NewLogger()
	c := &MotionController{drone: drone, leap: leap, logger: logger}
	return c
}

// Run function create the infinity loop to read all Frames until the user
// wants to land the drone
func (c *MotionController) Run() {

	drone := c.drone
	gestureManager := gesture.NewManager(drone)

	handManager := hand.NewManager(drone)
	currentHand := leap.Hand{}

	isProcessed := make(chan bool)

	c.leap.On(leap.GestureEvent, func(data interface{}) {

		go func() {
			g := data.(leap.Gesture)
			c.logger.ShowGesture(g)
			gestureManager.Process(g)

			isProcessed <- true
		}()

	})

	c.leap.On(leap.HandEvent, func(data interface{}) {

		go func() {

			isGesture := <-isProcessed
			if isGesture {
				fmt.Println("Is gesture before")
				return
			}

			isProcessed <- false
			currentHand = data.(leap.Hand)
			c.logger.ShowHand(currentHand)
			handManager.Process(currentHand)

		}()
	})
	/*
		c.leap.On(leap.MessageEvent, func(data interface{}) {
			f := data.(leap.Frame)
			isOpenHand := pointable.IsOpenHand(f.Pointables)
			fmt.Printf("Is open hand:%v \n", isOpenHand)
		})
	*/
}
