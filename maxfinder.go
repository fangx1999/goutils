package utils

import (
	"sync/atomic"
)

// ICompare is an interface for comparing objects.
// Objects implementing this interface must be able to compare themselves
// with other objects of the same type.
type ICompare interface {
	// LessThan returns true if the current object is less than the other object,
	// false otherwise.
	// The other object must be of the same type as the current object.
	LessThan(other ICompare) bool
}

// MaxFinder is a struct that finds the maximum value among a set of objects
type MaxFinder struct {
	// We use an atomic value to allow concurrent access to the maximum value
	maxValue atomic.Value
}

// NewMaxFinder creates a new MaxFinder with an initial maximum value
func NewMaxFinder(initMax ICompare) *MaxFinder {
	finder := &MaxFinder{}
	finder.maxValue.Store(initMax)
	return finder
}

// CompareAndSwap compares the given value with the current maximum value.
// If the given value is greater than the current maximum value, it replaces
// the current maximum value with the given value.
func (m *MaxFinder) CompareAndSwap(v ICompare) {
	for {
		oldMax := m.maxValue.Load().(ICompare)
		if v.LessThan(oldMax) {
			break
		}

		if m.maxValue.CompareAndSwap(oldMax, v) {
			break
		}
	}
}

// GetMax returns the current maximum value
func (m *MaxFinder) GetMax() ICompare {
	return m.maxValue.Load().(ICompare)
}
