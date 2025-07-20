package sleeper

import "time"

// Wrapper on original time package to make it testable!
type RealSleeper struct{}

func (s *RealSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
