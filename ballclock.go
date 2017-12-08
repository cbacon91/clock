package clock

import (
	"errors"
	"reflect"
	"time"
)

func main() {

}

type ballClock struct {
	main, minute, fiveMin, hour []int
}

func (c *ballClock) addMinute() {
	/*
		if minute arm currently has 4
			move balls in reverse order of addition to main //fn
			//minuteballs := reverseSlice(c.minute)
			c.minute.empty() / c.minute = []int
			c.main = append(c.main, minuteBalls)

			if 5min arm has currently has 11
				move balls in reverse order of addition to main
				if hour arm has currently has 11
					move balls in reverse order of addition to main
					add ball to main
				else
					append ball to hour arm
			else
				append ball to 5min arm
		else
			append ball to minute arm
	*/
}

// ReverseSlice takes a slice and returns a reversed copy of it.
func ReverseSlice(forward []int) []int {
	reverse := make([]int, len(forward))
	copy(reverse, forward)

	for front, back := 0, len(reverse)-1; front < back; front, back = front+1, back-1 {
		reverse[front], reverse[back] = reverse[back], reverse[front]
	}
	return reverse
}

// CycleDays takes a number of balls and determines the number of 24-hour periods which elapse before the clock returns to its initial ordering.
// It returns the number of days taken to return to inital order, the time in seconds it took to run, and potential errors.
func CycleDays(ballCount int) (int, float64, error) {
	start := time.Now()

	if ballCount < 27 || ballCount > 127 {
		return 0, 0, errors.New("can only run with ballCount between 27 and 127")
	}

	clock := newClock(ballCount)

	// setup the initial stack structure to compare against
	initialQ := make([]int, ballCount)
	copy(initialQ, clock.main)

	//run first case before this, otherwise we have an infinite loop here
	for ok := true; ok; ok = !reflect.DeepEqual(initialQ, clock.main) {
		clock.addMinute()
	}

	elapsed := time.Since(start).Seconds()
	return 0, elapsed, nil
}

func newClock(ballCount int) ballClock {
	clock := ballClock{}

	for i := 1; i <= ballCount; i++ {
		clock.main = append(clock.main, i)
	}

	return clock
}
