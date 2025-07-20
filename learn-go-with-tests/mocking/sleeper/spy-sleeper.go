package sleeper

const (
	write = "write"
	sleep = "sleep"
)

// Test Sleeper object against Sleeper interface!
type SpySleeper struct {
	Calls []string
}

func (s *SpySleeper) Sleep() {
	// s.Calls++ // helps in tracking no. of calls in tests
	s.Calls = append(s.Calls, sleep)
}

// Mock implementation for io.Writer Write() method
func (s *SpySleeper) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}
