package scheduler

import "youtube/interfaces"

type SimpleScheduler struct {
	in chan<- interfaces.Task
}

func (s *SimpleScheduler) SetChannel(in chan<- interfaces.Task) {
	s.in = in
}

func NewSimpleScheduler(in chan<- interfaces.Task) *SimpleScheduler {
	return &SimpleScheduler{in: in}
}

func (s *SimpleScheduler) Submit(task interfaces.Task) {
	go func() {
		s.in <- task
	}()
}
