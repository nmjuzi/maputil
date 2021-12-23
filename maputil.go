package maputil

import (
	"sync"
)

type UniqMap struct {
	sync.RWMutex
	m map[string]bool
}

func NewUniqMap(size int) *UniqMap {
	if size > 0 {
		return &UniqMap{m: make(map[string]bool, size)}
	} else {
		return &UniqMap{m: make(map[string]bool)}
	}
}

func (s *UniqMap) AddNotContains(key string) bool {
	s.Lock()
	defer s.Unlock()
	if s.m[key] {
		return false
	} else {
		s.m[key] = true
		return true
	}
}
