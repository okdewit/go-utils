package timeutils

import "time"

func QuarterBoundaries(y int, quarter int) (a time.Time, b time.Time) {
	return time.Date(y, time.Month(quarter*(quarter-1)+1), 1, 0, 0, 0, 0, time.UTC),
		time.Date(y,time.Month(quarter*(quarter-1)+4), 1, 0, 0, 0, -1, time.UTC)
}
