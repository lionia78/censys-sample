package kvstore

import (
	"sync"
)

type InMemoryStore struct {
	// mutex - to support multi go-routines (multithreaded) safety
	mutex sync.RWMutex

	// store - actual store (map)
	store map[string]string
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		store: make(map[string]string),
	}
}

func (s *InMemoryStore) Put(key string, value string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.store[key] = value
}

func (s *InMemoryStore) Get(key string) (string, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	v, ok := s.store[key]
	return v, ok
}

func (s *InMemoryStore) Delete(key string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, ok := s.store[key]; !ok {
		return false
	}
	
	delete(s.store, key)
	return true
}
