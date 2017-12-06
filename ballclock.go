package clock

import (
	"errors"
	"time"
)

func main() {

}

type ball int
type ballClock struct {
	main          []ball
	minute        [4]ball
	fiveMin, hour [11]ball
}

// CycleDays takes a number of balls and determines the number of 24-hour periods which elapse before the clock returns to its initial ordering.
// It returns the number of days taken to return to inital order, the time in seconds it took to run, and potential errors.
func CycleDays(ballCount int) (int, float64, error) {
	start := time.Now()

	if ballCount < 27 || ballCount > 127 {
		return 0, 0, errors.New("can only run with ballCount between 27 and 127")
	}

	clock := ballClock{
		make([]ball, ballCount),
		[4]ball{},
		[11]ball{},
		[11]ball{},
	}

	for i := range clock.main {
		clock.main[i] = ball(i)
	}

	// setup the initial queue structure
	initialQ := make([]ball, ballCount)
	copy(initialQ, clock.main)

	//run first case before this, otherwise we have an infinite loop here
	for !IsSameQueue(initialQ, clock.main) {
		//run logic while it's not the same queue

		//this is going to be logic shared between clock functions
		//handle minute, fiveminute, etc
		//method on ballclock?
	}

	elapsed := time.Since(start).Seconds()
	return 0, elapsed, nil
}

// IsSameQueue compares two slices of ball to verify same values and same ordering
func IsSameQueue(q1, q2 []ball) bool {
	if len(q1) != len(q2) {
		return false
	}

	for i, v := range q1 {
		if v != q2[i] {
			return false
		}
	}

	return true
}
