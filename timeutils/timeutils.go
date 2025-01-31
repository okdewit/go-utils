package timeutils

import (
	"fmt"
	"strconv"
	"time"
)

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

type Quarter struct {
	Y int;
	Q int;
}


func ToQuarter(t time.Time) (Quarter) {
	return Quarter{
		Y: t.Year(),
		Q: int((time.Now().Month()+2)/3),
	}
}

func (quarter *Quarter) String() string {
	return strconv.Itoa(quarter.Y) + "-Q" + strconv.Itoa(quarter.Q)
}

func (quarter Quarter) IsCurrent() bool {
	current := ToQuarter(time.Now())
	return quarter.Y == current.Y && quarter.Q == current.Q
}

func (quarter Quarter) IsFuture() bool {
	current := ToQuarter(time.Now())
	return quarter.Y > current.Y || (quarter.Y == current.Y && quarter.Q > current.Q)
}

func (quarter Quarter) IsPast() bool {
	current := ToQuarter(time.Now())
	return quarter.Y < current.Y || (quarter.Y == current.Y && quarter.Q < current.Q)
}

func (quarter *Quarter) Next() *Quarter {
	if quarter.Q == 4 {quarter.Y++}
	quarter.Q = (quarter.Q % 4) + 1
	return quarter
}

func (quarter *Quarter) Previous() *Quarter {
	if quarter.Q == 1 {quarter.Y--}
	// Modulo is not same as remainder for negatives
	quarter.Q = ((quarter.Q + 2) % 4) + 1
	return quarter
}

func (quarter *Quarter) Boundaries() (a time.Time, b time.Time) {
	return time.Date(quarter.Y, time.Month((quarter.Q-1)*3+1), 1, 0, 0, 0, 0, time.UTC),
		time.Date(quarter.Y,time.Month((quarter.Q-1)*3+4), 1, 0, 0, 0, -1, time.UTC)
}


