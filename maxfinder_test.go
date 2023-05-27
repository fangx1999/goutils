package utils

import (
	"math/rand"
	"sync"
	"testing"
)

type MyInt int

func (i MyInt) LessThan(other ICompare) bool {
	return i < other.(MyInt)
}

func findMaximumLinearly(myInts []MyInt) MyInt {
	if len(myInts) <= 0 {
		panic("slice is nil or empty")
	}
	max := myInts[0]
	for i := 1; i < len(myInts); i++ {
		if myInts[i] > max {
			max = myInts[i]
		}
	}
	return max
}

func generateRandomInts(size int) []MyInt {
	ints := make([]MyInt, size)
	for i := 0; i < size; i++ {
		ints[i] = MyInt(rand.Int())
	}
	return ints
}

func TestFindMaximum(t *testing.T) {
	ints := generateRandomInts(1000000)
	max := findMaximumLinearly(ints)

	finder := NewMaxFinder(MyInt(0))
	ch := make(chan struct{}, 100)
	var wg sync.WaitGroup

	for i := 0; i < len(ints); i++ {
		wg.Add(1)
		v := ints[i]
		go func() {
			defer wg.Done()
			ch <- struct{}{}
			finder.CompareAndSwap(v)
			<-ch
		}()
	}

	wg.Wait()

	if finder.GetMax().(MyInt) != max {
		t.Errorf("Expected %d, but got %d", max, finder.GetMax().(MyInt))
	}
}
