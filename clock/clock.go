// Package clock implement a simple h:m clock
package clock

import "fmt"

// A clock is a data wich contain hours and minutes
type Clock struct {
	hours   int
	minutes int
}

// Add minutes to clock
func (cl Clock) Add(minutes int) Clock {
	if minutes < 0 {
		return cl.Subtract(-minutes)
	}
	for cl.minutes+minutes >= 60 {
		cl.hours++
		minutes -= 60
		if cl.hours > 23 {
			cl.hours = 0
		}
	}
	cl.minutes += minutes
	return cl
}

// Subtract minutes to clock
func (cl Clock) Subtract(minutes int) Clock {
	if minutes < 0 {
		return cl.Add(-minutes)
	}
	for cl.minutes-minutes < 0 {

		cl.hours--
		minutes -= 60
		if cl.hours < 0 {
			cl.hours = 23
		}
	}
	cl.minutes -= minutes
	return cl
}

// New function create a clock from given hour and minutes value
func New(hours, minutes int) Clock {
	cl := Clock{0, 0}
	return cl.Add(minutes + (hours * 60))
}

// String return clock formatted to be human readable
func (cl Clock) String() string {
	return fmt.Sprintf("%.2d:%.2d", cl.hours, cl.minutes)
}
