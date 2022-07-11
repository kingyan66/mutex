package safequeue

import "sync"

type SafeQueue struct {
	data []interface{}
	mu   sync.Mutex
}

func NewSafeQueue(n int) *SafeQueue {
	return &SafeQueue{data: make([]interface{}, 0, n)}
}

func (s *SafeQueue) Enqueue(data interface{}) {
	s.mu.Lock()
	s.data = append(s.data, data)
	s.mu.Unlock()
}

func (s *SafeQueue) Dequeue() interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.data) == 0 {
		return nil
	}
	v := s.data[0]
	s.data = s.data[1:]
	return v
}
