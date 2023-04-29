package stopwatch

import "time"

type Stopwatch struct {
	st  time.Time
	dur []time.Duration
}

func (s *Stopwatch) Start() {
	s.st = time.Now()
	s.dur = s.dur[:0]
}

func (s *Stopwatch) SaveSplit() {
	s.dur = append(s.dur, time.Now().Sub(s.st))
}

func (s *Stopwatch) GetResults() []time.Duration {
	return s.dur
}
