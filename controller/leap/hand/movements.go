package hand

import "math"

// IsXMovement  verifies if there is a shift in the X axis
//and if it is up to the defined threshold.
func IsXMovement(shift float64) bool {
	return math.Abs(shift) > DirectionThreshold
}
