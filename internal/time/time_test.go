package time_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	itime "github.com/ivanlemeshev/n/internal/time"
)

func TestNewRealClock(t *testing.T) {
	clock := itime.NewRealClock()

	require.Equal(t,
		time.Now().Format("2006-01-02T15:04:05"),
		clock.Now().Format("2006-01-02T15:04:05"))
}

func TestNewMockClock(t *testing.T) {
	now := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)

	clock := itime.NewMockClock(func() time.Time {
		return now
	})

	require.Equal(t, now, clock.Now())
}
