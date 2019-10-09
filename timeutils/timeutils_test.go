package timeutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestQuarterBoundaries(t *testing.T) {
	t.Run("Q1", func(t *testing.T) {
		s, e := (&Quarter{Y: 2010, Q: 1}).Boundaries()
		assert.Equal(t, "2010-01-01", s.Format("2006-01-02"))
		assert.Equal(t, "2010-03-31", e.Format("2006-01-02"))
	})

	t.Run("Q2", func(t *testing.T) {
		s, e := (&Quarter{Y: 1900, Q: 2}).Boundaries()
		assert.Equal(t, "1900-04-01", s.Format("2006-01-02"))
		assert.Equal(t, "1900-06-30", e.Format("2006-01-02"))
	})

	t.Run("Q3", func(t *testing.T) {
		s, e := (&Quarter{Y: 2021, Q: 3}).Boundaries()
		assert.Equal(t, "2021-07-01", s.Format("2006-01-02"))
		assert.Equal(t, "2021-09-30", e.Format("2006-01-02"))
	})

	t.Run("Q4", func(t *testing.T) {
		s, e := (&Quarter{Y: 1501, Q: 4}).Boundaries()
		assert.Equal(t, "1501-10-01", s.Format("2006-01-02"))
		assert.Equal(t, "1501-12-31", e.Format("2006-01-02"))
	})
}

func TestFmtDuration(t *testing.T) {
	assert.Equal(t, "00:12", FmtDuration(12 * time.Minute))
	assert.Equal(t, "1 days 02:00", FmtDuration(26 * time.Hour))
	assert.Equal(t, "100 days 00:00", FmtDuration(100 * 24 * time.Hour))
}
