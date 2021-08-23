package main

import (
	"fmt"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/leap"
	controller "tello/controller/leap"
)

func main() {
	leapMotionAdaptor := leap.NewAdaptor("127.0.0.1:6437")
	l := leap.NewDriver(leapMotionAdaptor)

	drone := tello.NewDriver("8888")
	fmt.Println(drone)
	c := controller.NewLeapMotion(l)

	work := func() {
		c.Run()
	}

	robot := gobot.NewRobot("tello-leap",
		[]gobot.Connection{leapMotionAdaptor},
		[]gobot.Device{l},
		work,
	)

	robot.Start()

	defer robot.Stop()
}
