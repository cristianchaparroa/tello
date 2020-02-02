package main

import (
	controller "tello/controller/leap"

	"gobot.io/x/gobot/platforms/leap"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

func main() {
	leapMotionAdaptor := leap.NewAdaptor("127.0.0.1:6437")
	l := leap.NewDriver(leapMotionAdaptor)
	drone := tello.NewDriver("8888")

	c := controller.NewController(drone, l)

	work := func() {
		c.Run()
	}

	robot := gobot.NewRobot("tello-leap",
		[]gobot.Connection{leapMotionAdaptor},
		[]gobot.Device{l, drone},
		work,
	)

	robot.Start()

	defer robot.Stop()
}
