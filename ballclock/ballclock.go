package ballclock

import (
	"encoding/json"
	"errors"
	"reflect"
	"time"
)

// ErrInvalidBallCount is an error thrown if the number of balls supplied is not between 27 and 127.
var ErrInvalidBallCount = errors.New("can only run with ballCount between 27 and 127")
// ErrMinutesNotSpecified is an error thrown if the number of minutes to run for is not specified for clock mode 2.
var ErrMinutesNotSpecified = errors.New("must specify positive number of minutes to run for")

// BallClock is the structure of the ball clock. It has four 'tracks' to represent timekeeping -
// Main (the deposit / unused balls), Min (representing a single minute), FiveMin (represent batches
// of five minute intervals), Hour (representing hour increments), and an internal property
// representing full completions of the clock to know the number of twelve-hour periods or full
// rotations of the clock have occurred.
type BallClock struct {
	twelveHours              int
	Main, Min, FiveMin, Hour []int
}

func (c *BallClock) String() string {
	jason, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	return string(jason)
}

// AddMinute adds a minute to the ball clock. Attempts to move a ball from main to Min.
// If Min is full, it appends Min to main in reverse order and attempts to move the
// ball to fivemin instead. If fivemin is full, it repeats this process with hour. Finally,
// if hour is full the clock is reset
func (c *BallClock) AddMinute () bool {
	//remove from main
	newBall := c.Main[0] // take the least recently used ball
	c.Main = c.Main[1:]

	if len(c.Min) == 4 {

		minuteBalls := ReverseSlice(c.Min)
		c.Min = make([]int, 0)
		c.Main = append(c.Main, minuteBalls...)

		if len(c.FiveMin) == 11 {
			fiveMinBalls := ReverseSlice(c.FiveMin)
			c.FiveMin = make([]int, 0)
			c.Main = append(c.Main, fiveMinBalls...)

			if len(c.Hour) == 11 {
				hourBalls := ReverseSlice(c.Hour)
				c.Hour = make([]int, 0)
				c.Main = append(c.Main, hourBalls...)
				c.Main = append(c.Main, newBall)
				c.twelveHours++
				return true
			} 

			c.Hour = append(c.Hour, newBall)
		} else {
			c.FiveMin = append(c.FiveMin, newBall)
		}
	} else {
		c.Min = append(c.Min, newBall)
	}
	return false
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
func CycleDays(ballCount int) (float32, float64, error) {
	start := time.Now()

	if ballCount < 27 || ballCount > 127 {
		return 0, 0, ErrInvalidBallCount
	}

	clock := newClock(ballCount)

	// setup the initial stack structure to compare against
	initialQ := make([]int, ballCount)
	copy(initialQ, clock.Main)

	for ok, isMainTrackFull := true, true; ok; ok = !isMainTrackFull || !reflect.DeepEqual(initialQ, clock.Main) {
		isMainTrackFull = clock.AddMinute()
	}

	elapsed := time.Since(start).Seconds()
	return float32(clock.twelveHours) / 2, elapsed, nil
}

// Clock is the clock mode.
func Clock(ballCount, runToMinutes int) (BallClock, float64, error) {
	start := time.Now()
	var clock BallClock

	if ballCount < 27 || ballCount > 127 {
		return clock, 0, ErrInvalidBallCount
	}
	if runToMinutes < 1 {
		return clock, 0, ErrMinutesNotSpecified
	}

	clock = newClock(ballCount)

	//run first case before this, otherwise we have an infinite loop here
	for i := 1; i <= runToMinutes; i++ {
		clock.AddMinute()
	}

	elapsed := time.Since(start).Seconds()
	return clock, elapsed, nil
}

func newClock(ballCount int) BallClock {
	clock := BallClock{}

	for i := 1; i <= ballCount; i++ {
		clock.Main = append(clock.Main, i)
	}

	return clock
}
