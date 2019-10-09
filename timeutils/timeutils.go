package timeutils

import (
	"fmt"
	"time"
)

func QuarterBoundaries(y int, quarter int) (a time.Time, b time.Time) {
	return time.Date(y, time.Month((quarter-1)*3+1), 1, 0, 0, 0, 0, time.UTC),
		time.Date(y,time.Month((quarter-1)*3+4), 1, 0, 0, 0, -1, time.UTC)
}

func FmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	x := d / (time.Hour * 24)
	d -= x * (time.Hour * 24)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	if x > 0 {
		return fmt.Sprintf("%d days %02d:%02d", x, h, m)
	}
	return fmt.Sprintf("%02d:%02d", h, m)
}
