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
		days, time, err := CycleDays(c.in)
		if days != c.days {
			t.Errorf("CycleDays(%v) == %v, want %v (days)", c.in, days, c.days)
		}
		if time != c.time {
			t.Errorf("CycleDays(%v) == %v, want %v (time)", c.in, time, c.time)
		}
		if err != nil && c.err != "" && err.Error() != c.err {
			t.Errorf("CycleDays(%v) == %v, want %v (err)", c.in, err.Error(), c.err)
		}
	}
}
