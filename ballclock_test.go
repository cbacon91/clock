package clock

import (
	"errors"
	"reflect"
	"testing"
)

func TestCycleDays(t *testing.T) {
	cases := []struct {
		in   int
		days float32
		err  error
	}{
		{26, 0, errors.New("can only run with ballCount between 27 and 127")},
		{27, 23, nil},
		{30, 15, nil},
		{45, 378, nil},
		{128, 0, errors.New("can only run with ballCount between 27 and 127")},
	}
	for _, c := range cases {
		days, _, err := CycleDays(c.in)
		if days != c.days {
			t.Errorf("CycleDays(%v) == %v, want %v (days)", c.in, days, c.days)
		}
		if !reflect.DeepEqual(err, c.err) { //!= nil && c.err != nil && err.Error() != c.err.Error() {
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
		{26, 2, BallClock{}, errors.New("can only run with ballCount between 27 and 127")},
		{128, 2, BallClock{}, errors.New("can only run with ballCount between 27 and 127")},
		{30, 0, BallClock{}, errors.New("must specify positive number of minutes to run for")},
		{30, -100, BallClock{}, errors.New("must specify positive number of minutes to run for")},
		{30, 325, BallClock{0, []int{11, 5, 26, 18, 2, 30, 19, 8, 24, 10, 29, 20, 16, 21, 28, 1, 23, 14, 27, 9}, []int{}, []int{22, 13, 25, 3, 7}, []int{6, 12, 17, 4, 15}}, nil},
	}
	for _, c := range cases {
		clock, _, err := Clock(c.balls, c.minutes)
		if !reflect.DeepEqual(clock, c.expected) {
			t.Errorf("Clock(%v, %v) == %v, want %v", c.balls, c.minutes, clock, c.expected)
		}
		if !reflect.DeepEqual(err, c.err) { //!= nil && c.err != nil && err.Error() != c.err.Error() {
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
			t.Errorf("ReverseSlice(%v) == %v, got %v ", c.in, c.expected, result)
		}
	}
}
