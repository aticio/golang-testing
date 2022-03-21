package sort

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortIncreasingOrder(t *testing.T) {
	// init
	elements := GetElements(10)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 9, elements[0])
	assert.EqualValues(t, 0, elements[len(elements)-1])

	timeoutChan := make(chan bool, 1)
	defer close(timeoutChan)

	go func() {
		// execution
		BubbleSort(elements)
		timeoutChan <- false
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		timeoutChan <- true
	}()

	if <-timeoutChan {
		assert.Fail(t, "bubble sort took more than 50 ms")
		return
	}

	// validation
	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 9, elements[len(elements)-1])
}

func TestSortIncreasingOrder(t *testing.T) {
	// init
	elements := GetElements(10)

	// execution
	Sort(elements)

	// validation
	assert.EqualValues(t, 0, elements[0], "first element should be 0")
	assert.EqualValues(t, 9, elements[len(elements)-1], "last element should be 9")
}

func BenchmarkBubbleSort(b *testing.B) {
	// init
	elements := GetElements(100000)

	// execution
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	// init
	elements := GetElements(100000)

	// execution
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
