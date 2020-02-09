package frame

import "gobot.io/x/gobot/platforms/leap"

// It verifies if all pointables(fingers) are open
// to determines that hand is open.
func IsOpenHand(ps []leap.Pointable) bool {

	counter := 0
	for _, p := range ps {
		if p.Extended {
			counter++
		}
	}

	return  counter > 0
}