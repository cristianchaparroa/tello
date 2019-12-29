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
