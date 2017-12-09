package clock

import (
	"encoding/json"
	"errors"
	"reflect"
	"time"
)

func main() {

}

// BallClock is a thing
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
func (c *BallClock) AddMinute() {
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
			} else {
				c.Hour = append(c.Hour, newBall)
			}
		} else {
			c.FiveMin = append(c.FiveMin, newBall)
		}
	} else {
		c.Min = append(c.Min, newBall)
	}
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
		return 0, 0, errors.New("can only run with ballCount between 27 and 127")
	}

	clock := newClock(ballCount)

	// setup the initial stack structure to compare against
	initialQ := make([]int, ballCount)
	copy(initialQ, clock.Main)

	//run first case before this, otherwise we have an infinite loop here
	for ok := true; ok; ok = !reflect.DeepEqual(initialQ, clock.Main) {
		clock.AddMinute()
	}

	elapsed := time.Since(start).Seconds()
	return float32(clock.twelveHours) / 2, elapsed, nil
}

// Clock is the clock mode.
func Clock(ballCount, runToMinutes int) (BallClock, float64, error) {
	start := time.Now()
	var clock BallClock

	if ballCount < 27 || ballCount > 127 {
		return clock, 0, errors.New("can only run with ballCount between 27 and 127")
	}
	if runToMinutes < 1 {
		return clock, 0, errors.New("must specify positive number of minutes to run for")
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
