package leap

import (
	"fmt"

	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
)

// Controller is a controller in charge to handle the
// moviements of tello drone with Controller device.
type Controller struct {
	drone    *tello.Driver
	leap     *leap.Driver
	listener MotionListener
	logger   Logger
}

// NewController generates pointer to Controller
func NewController(drone *tello.Driver, leap *leap.Driver) *Controller {
	listener := NewleapMotionListener()
	logger := NewLogger()
	return &Controller{drone: drone, leap: leap, listener: listener, logger: logger}
}

// Run function create the infinity loop to read all Frames until the user
// wants to land the drone
func (c *Controller) Run() {

	c.leap.On(leap.GestureEvent, func(data interface{}) {
		fmt.Println("--> Gesture event")

		gesture := data.(leap.Gesture)
		c.logger.ShowGesture(gesture)
		isTakeOff := c.listener.IsTakeOffEvent(gesture)

		if isTakeOff {
			fmt.Println("-- Is take off")
			c.TakeOff()
		}

		fmt.Println("<-- Gesture event")
	})

}

// TakeOff the drone with the
func (c *Controller) TakeOff() {
	c.drone.TakeOff()
}

// Up  shift the drone up
func (c *Controller) Up() {

}

// Down  shift the drone down
func (c *Controller) Down() {}

// Right shift the drone right
func (c *Controller) Right() {}

// Left  shift the drone  left.
func (c *Controller) Left() {}

// Land the drone
func (c *Controller) Land() {}
