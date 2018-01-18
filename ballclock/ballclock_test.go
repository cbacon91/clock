package ballclock

import (
	"reflect"
	"testing"
)

func TestCycleDays(t *testing.T) {
	cases := []struct {
		in   int
		days float32
		err  error
	}{
		{26, 0, ErrInvalidBallCount},
		{27, 23, nil},
		{30, 15, nil},
		{45, 378, nil},
		{128, 0, ErrInvalidBallCount},
		{123, 108855, nil}, // basically worst case scenario
	}
	for _, c := range cases {
		days, _, err := CycleDays(c.in)
		if days != c.days {
			t.Errorf("CycleDays(%v) == %v, want %v (days)", c.in, days, c.days)
		}
		if !reflect.DeepEqual(err, c.err) {
			t.Errorf("CycleDays(%v) == %v, want %v (err)", c.in, err.Error(), c.err)
		}
	}
}

func TestClock(t *testing.T) {
	cases := []struct {
		balls, minutes int
		expected       BallClock
		err            error
	}{
		{26, 2, BallClock{}, ErrInvalidBallCount},
		{128, 2, BallClock{}, ErrInvalidBallCount},
		{30, 0, BallClock{}, ErrMinutesNotSpecified},
		{30, -100, BallClock{}, ErrMinutesNotSpecified},
		{30, 325, BallClock{0, []int{11, 5, 26, 18, 2, 30, 19, 8, 24, 10, 29, 20, 16, 21, 28, 1, 23, 14, 27, 9}, []int{}, []int{22, 13, 25, 3, 7}, []int{6, 12, 17, 4, 15}}, nil},
	}
	for _, c := range cases {
		clock, _, err := Clock(c.balls, c.minutes)
		if !reflect.DeepEqual(clock, c.expected) {
			t.Errorf("Clock(%v, %v) == %v, want %v", c.balls, c.minutes, clock, c.expected)
		}
		if !reflect.DeepEqual(err, c.err) {
			t.Errorf("Clock(%v, %v) == %v, want %v", c.balls, c.minutes, err, c.err)
		}
	}
}

func TestReverseSlice(t *testing.T) {
	cases := []struct {
		in, expected []int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{2, 1, 3}, []int{3, 1, 2}},
		{[]int{-1, 2, -3}, []int{-3, 2, -1}},
	}
	for _, c := range cases {
		result := ReverseSlice(c.in)
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("ReverseSlice(%v) == %v, want %v ", c.in, result, c.expected)
		}
	}
}

func TestAddMinute(t *testing.T) {
	cases := []struct {
		inClock, outClock BallClock
	}{
		{ // add minute
			BallClock{0, []int{1}, []int{}, []int{}, []int{}},
			BallClock{0, []int{}, []int{1}, []int{}, []int{}},
		},
		{ // add minute - rollover to 5
			BallClock{0, []int{1}, []int{2, 3, 4, 5}, []int{}, []int{}},
			BallClock{0, []int{5, 4, 3, 2}, []int{}, []int{1}, []int{}},
		},
		{ // add minute - rollover to hr
			BallClock{0, []int{1}, []int{2, 3, 4, 5}, []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, []int{}},
			BallClock{0, []int{5, 4, 3, 2, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6}, []int{}, []int{}, []int{1}},
		},
		{ // add minute - rollover to 12hour
			BallClock{0, []int{1}, []int{2, 3, 4, 5}, []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, []int{17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27}},
			BallClock{1, []int{5, 4, 3, 2, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 1}, []int{}, []int{}, []int{}},
		},
	}

	for _, c := range cases {
		c.inClock.AddMinute()
		if !reflect.DeepEqual(c.inClock, c.outClock) {
			t.Errorf("BallClock.AddMinute() == %v, wanted %v ", c.inClock, c.outClock)
		}
	}
}
