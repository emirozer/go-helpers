package helpers

import (
	rnd "math/rand"
	"time"
)

// RandInt returns a random integer between the min and max you set
func RandInt(min int, max int) int {
	rnd.Seed(time.Now().UnixNano())
	return min + rnd.Intn(max-min)
}
