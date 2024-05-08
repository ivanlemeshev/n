package time

import "time"

// Clock is an interface for getting the current time.
type Clock interface {
	Now() time.Time
}

// RealClock is a Clock that returns the actual current time.
type RealClock struct{}

// NewRealClock returns a new RealClock.
func NewRealClock() *RealClock {
	return &RealClock{}
}

// Now returns the current time.
func (RealClock) Now() time.Time {
	return time.Now()
}

// MockClock is a Clock that returns a fixed time for testing.
type MockClock struct {
	NowFn func() time.Time
}

// NewMockClock returns a new MockClock.
func NewMockClock(nowFn func() time.Time) *MockClock {
	return &MockClock{NowFn: nowFn}
}

// Now returns the fixed time.
func (m MockClock) Now() time.Time {
	return m.NowFn()
}
