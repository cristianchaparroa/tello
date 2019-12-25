package controller

import "gobot.io/x/gobot"

// Controller provides the function to
// manipulate a drone
type Controller interface {
	Up()
	Down()
	Right()
	Left()
	TakeOff()
	Land()
	Run()

	GetDevice() gobot.Device
	GetEvents() []*EventHandler
}
