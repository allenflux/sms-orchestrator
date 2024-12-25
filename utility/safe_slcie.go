package utility

import (
	"fmt"
	"sync"
)

type SafeSlice struct {
	mu    sync.Mutex
	slice []interface{}
}

func NewSafeSlice() *SafeSlice {
	return &SafeSlice{}
}

func (s *SafeSlice) Get(index int) (interface{}, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if index < 0 || index >= len(s.slice) {
		return nil, fmt.Errorf("index out of bounds")
	}
	return s.slice[index], nil
}

func (s *SafeSlice) Append(value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.slice = append(s.slice, value)
}

func (s *SafeSlice) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.slice)
}
