package leap

import (
	"fmt"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
	"time"
)

// MotionController is in charge to capture the moviments of hands throught
// the leap motion device and send commands to Drone
type MotionController struct {
	drone *tello.Driver
	leap  *leap.Driver
}

// NewLeapMotion creates a pointer to MotionController
func NewLeapMotion(l *leap.Driver) *MotionController {
	return &MotionController{leap: l}
}

// Run creates a infinity loop reading all frames generated by user throughout
// leap motion until user wants to land the drone.
func (c *MotionController) Run() {
	events := c.leap.Subscribe()
	for {
		select {
		case e := <-events:
			eventName := e.Name

			if leap.GestureEvent == eventName {
				fmt.Println("it is a gesture")
			}

			if leap.HandEvent == eventName {
				fmt.Println("It is a hand")
			}

			// All the time by default is a message
			// so it should be careful that is the message required
			if leap.MessageEvent == eventName {
				fmt.Println("It is a message")
			}
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}

}
