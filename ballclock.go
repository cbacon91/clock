package clock

import "errors"

func main() {

}

// CycleDays ...
func CycleDays(ballCount int) (int, float64, error) {
	if ballCount < 27 || ballCount > 127 {
		return 0, 0, errors.New("can only run with ballCount between 27 and 127")
	}

	return 0, 0, nil
}
