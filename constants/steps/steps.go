package steps

import "math"

// ALL --> step value for migrate all migrations using value of math.MaxInt64
var ALL = math.MaxInt64

// IsAll --> check whether the step value is all or not
func IsAll(stepsValue int) bool {
	if stepsValue == ALL {
		return true
	}

	return false
}
