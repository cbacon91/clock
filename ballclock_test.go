package clock

import (
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

func TestIsSameQueue(t *testing.T) {
	cases := []struct {
		q1, q2 []ball
		result bool
	}{
		{[]ball{1, 2, 3}, []ball{1, 2, 3}, true},  // success
		{[]ball{1, 2, 3}, []ball{3, 2, 1}, false}, // wrong order
		{[]ball{1, 2, 3}, []ball{2, 2, 3}, false}, // not same elements
		{[]ball{1, 2, 3}, []ball{2, 3}, false},    // q1 longer
		{[]ball{2, 3}, []ball{1, 2, 3}, false},    // q2 longer
	}
	for _, c := range cases {
		result := IsSameQueue(c.q1, c.q2)
		if result != c.result {
			t.Errorf("IsSameQueue(%v, %v) == %v, want %v ", c.q1, c.q2, result, c.result)
		}
	}
}
