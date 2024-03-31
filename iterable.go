package main

import (
	"fmt"
	"sort"
	"sync"
)

type sliceStream[T any] struct {
	contents []T
	rw sync.RWMutex
}

func NewStream[T any](s []T) Stream[T] {
	return &sliceStream[T]{
		contents: s,
	}
}

func Map[T any, K any](s Stream[T], f func(T) K) Stream[K] {
	result := make([]K, 0, s.Length())
	for _, v := range s.Slice() {
		result = append(result, f(v))
	}
	return NewStream(result)
}


func (s *sliceStream[T]) Filter(f func(T) bool) Stream[T] {
	s.rw.RLock()
	defer s.rw.RUnlock()
	result := make([]T, 0, len(s.contents))
	for _, v := range s.contents {
		if(f(v)){
			result = append(result, v)
		}
	}
	return NewStream(result)
}

func (s *sliceStream[T]) Length() int {
	return len(s.contents)
}

func (s *sliceStream[T]) Slice() []T {
	return s.contents
}

func (s *sliceStream[T]) Next() (result T, err error) {
	s.rw.Lock()
	defer s.rw.Unlock()
	if len(s.contents) == 0 {
		err = fmt.Errorf("No more elements")
		return
	}
	v := s.contents[0]
	s.contents = s.contents[1:]
	return v, nil
}

func (s *sliceStream[T]) Sort(f func(T, T) bool) Stream[T]{
	s.rw.RLock()
	result := make([]T, 0, len(s.contents))
	for _, v := range s.contents {
		result = append(result, v)
	}
	s.rw.RUnlock()
	sort.Slice(result, func(i, j int) bool {
		return f(result[i], result[j])
	})
	return NewStream(result)
}

type Stream[T any] interface {
	Filter(func(T) bool) Stream[T]
	Sort(func(T, T) bool) Stream[T]
	Length() int
	Slice() []T
	Next() (T, error)
}
