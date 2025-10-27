package utils

import "time"

type TimeProvider interface {
	Now() time.Time
}

type RealClock struct{}

func (rc RealClock) Now() time.Time {
	return time.Now()
}

type MockClock struct {
	NowFunc func() time.Time
}

func (mc MockClock) Now() time.Time {
	if mc.NowFunc == nil {
		return time.Date(1999, time.October, 18, 0, 0, 0, 0, time.UTC)
	}

	return mc.NowFunc()
}
