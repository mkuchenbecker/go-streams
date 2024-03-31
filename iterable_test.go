package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	s := NewStream([]int{1, 2, 3, 4, 5})
	result := s.Filter(func(v int) bool {
		return v%2 == 0
	})
	if result.Length() != 2 {
		t.Errorf("Expected 2, got %d", result.Length())
	}
	item1, err1 := result.Next()
	assert.Nil(t, err1)
	assert.Equal(t, 2, item1)

	item2, err2 := result.Next()
	assert.Nil(t, err2)
	assert.Equal(t, 4, item2)

	_, err3 := result.Next()
	assert.NotNil(t, err3)
}

func TestMap(t *testing.T) {
	s := NewStream([]int{1, 2, 3, 4, 5})
	result := Map(s, func(v int) int {
		return v * 2
	})
	if result.Length() != 5 {
		t.Errorf("Expected 5, got %d", result.Length())
	}
	contents := result.Slice()
	assert.Equal(t, []int{2, 4, 6, 8, 10}, contents)

	item1, err1 := result.Next()
	assert.Nil(t, err1)
	assert.Equal(t, 2, item1)

	item2, err2 := result.Next()
	assert.Nil(t, err2)
	assert.Equal(t, 4, item2)

	item3, err3 := result.Next()
	assert.Nil(t, err3)
	assert.Equal(t, 6, item3)

	item4, err4 := result.Next()
	assert.Nil(t, err4)
	assert.Equal(t, 8, item4)

	item5, err5 := result.Next()
	assert.Nil(t, err5)
	assert.Equal(t, 10, item5)

	_, err6 := result.Next()
	assert.NotNil(t, err6)
}

func TestSort(t *testing.T) {
	s := NewStream([]int{3, 2, 5, 1, 4})
	result := s.Sort(LessThan)
	assert.Equal(t, 5, result.Length())
	contents := result.Slice()
	assert.Equal(t, []int{1, 2, 3, 4, 5}, contents)

	s2 := NewStream([]int{3, 2, 5, 1, 4})
	result2 := s2.Sort(GreaterThan)
	assert.Equal(t, []int{5, 4, 3, 2, 1}, result2.Slice())
	assert.Equal(t, 5, result2.Length())
}
