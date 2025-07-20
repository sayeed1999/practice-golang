package sleeper

// Test Sleeper object against Sleeper interface!
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++ // helps in tracking no. of calls in tests
}
