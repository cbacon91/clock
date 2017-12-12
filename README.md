# BallClock

A ball clock is a device which measures time by adding balls to tracks (one-minute, five-minute, and hour) to keep track of time. This project is a full simulation of a physical ball-clock written in Go. 

It has two modes: 

1) CyleDays
  The first mode takes a single parameter specifying the number of balls and reports the number of balls given in the input and the number of days (24-hour periods) which elapse before the clock returns to its initial ordering.

2) Clock
The second mode takes two parameters, the number of balls and the number of minutes to run for. If the number of minutes is specified, the clock must run to the number of minutes and report the state of the tracks at that point in a JSON format.
