package clock

import (
	"reflect"
	"testing"
)

func TestBallCount(t *testing.T) {
	cases := []struct {
		in, days int
		time     float64
		err      string
	}{
		{26, 0, 0, "can only run with ballCount between 27 and 127"},
		{27, 0, 0, ""},
		{28, 0, 0, ""},
		{126, 0, 0, ""},
		{127, 0, 0, ""},
		{128, 0, 0, "can only run with ballCount between 27 and 127"},
	}
	for _, c := range cases {
		days, _, err := CycleDays(c.in)
		if days != c.days {
			t.Errorf("CycleDays(%v) == %v, want %v (days)", c.in, days, c.days)
		}
		// no test case for time because it's variable by definition
		if err != nil && c.err != "" && err.Error() != c.err {
			t.Errorf("CycleDays(%v) == %v, want %v (err)", c.in, err.Error(), c.err)
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
