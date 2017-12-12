package main

import (
	"flag"
	"fmt"

	"github.com/cbacon91/clock/ballclock"
)

func main() {
	modePtr := flag.Int("mode", 1, "Mode of the clock: 1 = cycledays, 2 = clock. Defaults to Mode 1.")
	ballCountPtr := flag.Int("ballCount", 27, "Number of balls to enter into the clock. Defaults to 27.")
	runToMinutesPtr := flag.Int("runToMinutes", 60, "Number of minutes to run for. Only used for mode 2. Defaults 720 (twelve hours).")

	flag.Parse()

	if *modePtr == 1 {
		days, seconds, err := ballclock.CycleDays(*ballCountPtr)

		if err != nil {
			fmt.Println(err.Error())
			return // just get outta here
		}

		fmt.Printf("%v balls cycle after %v days.\n", *ballCountPtr, days)
		fmt.Printf("Completed in %0.1f miliseconds (%0.3f seconds)", seconds*1000, seconds)
	} else if *modePtr == 2 {
		endState, _, err := ballclock.Clock(*ballCountPtr, *runToMinutesPtr)

		if err != nil {
			fmt.Println(err.Error())
			return // just get outta here
		}

		fmt.Println(endState.String())
	} else {
		fmt.Printf("Mode %v not recognized.", *modePtr)
	}
}
