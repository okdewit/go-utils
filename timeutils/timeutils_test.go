package timeutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuarterBoundaries(t *testing.T) {
	t.Run("Q1", func(t *testing.T) {
		s, e := QuarterBoundaries(2010, 1)
		assert.Equal(t, "2010-01-01", s.Format("2006-01-02"))
		assert.Equal(t, "2010-03-31", e.Format("2006-01-02"))
	})

	t.Run("Q2", func(t *testing.T) {
		s, e := QuarterBoundaries(1900, 2)
		assert.Equal(t, "1900-04-01", s.Format("2006-01-02"))
		assert.Equal(t, "1900-06-30", e.Format("2006-01-02"))
	})

	t.Run("Q3", func(t *testing.T) {
		s, e := QuarterBoundaries(2021, 3)
		assert.Equal(t, "2021-07-01", s.Format("2006-01-02"))
		assert.Equal(t, "2021-09-30", e.Format("2006-01-02"))
	})

	t.Run("Q4", func(t *testing.T) {
		s, e := QuarterBoundaries(1501, 4)
		assert.Equal(t, "1501-10-01", s.Format("2006-01-02"))
		assert.Equal(t, "1501-12-31", e.Format("2006-01-02"))
	})
}
