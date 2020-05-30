# Tello

This library provides multiple controllers to use easy the DJ Tello drone using Golang.

## Controllers

The following are the supported devices:

- [x] Key board
- [ ] Leap motion device


### Keyboard

You can handle the drone using the default keyboard  configuration. The following is the setup.

![](./images/keyboard-configuration.png)

#### Example

```go
package main

import (
	"tello/controller/keyboard"

	"gobot.io/x/gobot"
	tello "gobot.io/x/gobot/platforms/dji/tello"
)

func main() {
	drone := tello.NewDriver("8888")
	c := keyboard.NewController(drone)

	work := func() {
		c.Run()
	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{c.GetKeyBoard(), drone},
		work,
	)

	robot.Start()
}

```

### Leap Motion controller

You can use handle the drone using the Leap motion controller



```go
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

	c := controller.NewLeapMotion(drone, l)

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

```
