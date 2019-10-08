package timeutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuarterBoundaries(t *testing.T) {
	s, e := QuarterBoundaries(2019, 3)
	assert.Equal(t, "2019-07-01", s.Format("2006-01-02"))
	assert.Equal(t, "2019-09-30", e.Format("2006-01-02"))

	s, e = QuarterBoundaries(2010, 1)
	assert.Equal(t, "2010-01-01", s.Format("2006-01-02"))
	assert.Equal(t, "2010-03-31", e.Format("2006-01-02"))
}
