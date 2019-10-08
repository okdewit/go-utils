package timeutils

import "time"

func QuarterBoundaries(y int, quarter int) (a time.Time, b time.Time) {
	return time.Date(y, time.Month((quarter-1)*3+1), 1, 0, 0, 0, 0, time.UTC),
		time.Date(y,time.Month((quarter-1)*3+4), 1, 0, 0, 0, -1, time.UTC)
}
