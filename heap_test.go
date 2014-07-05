package heap_test

import (
	"github.com/aaylward/heap"
	"testing"
)

func Test_ExtractMinEmpty(t *testing.T) {
	minHeap := heap.NewMinHeap()
	_, err := minHeap.ExtractMin()

	if err == nil {
		t.Error("heap.ExtractMin should return an error on empty heaps!")
	} else {
		t.Log("Woo! Error returned.")
	}
}

func Test_InsertToHeap(t *testing.T) {
	minHeap := heap.NewMinHeap()
	minHeap.Insert(5)

	if minHeap.Size() != 1 {
		t.Error("Failed to insert an element!")
	} else {
		t.Log("Successfully inserted an element!")
	}
}

func Test_ExtractMin(t *testing.T) {
	minHeap := heap.NewMinHeap()
	minHeap.Insert(1)
	minHeap.Insert(2)
	minHeap.Insert(3)
	minHeap.Insert(-3)
	minHeap.Insert(12)

	if minHeap.Size() != 5 {
		t.Error("Heap is the wrong size!")
	}

	min, err := minHeap.ExtractMin()

	if min != -3 || err != nil {
		t.Error("Failed to ExtractMin!")
	}

	if minHeap.Size() != 4 {
		t.Error("Failed to remove the minimum element!")
	}
}

func Test_Peek(t *testing.T) {

	minHeap := heap.NewMinHeap()
	minHeap.Insert(6)
	minHeap.Insert(-12)
	minHeap.Insert(3)
	minHeap.Insert(-3)
	minHeap.Insert(12)

	if minHeap.Size() != 5 {
		t.Error("Heap is the wrong size!")
	}

	if minHeap.Peek() != -12 {
		t.Error("Peek() is not the minumum element!")
	}

	if minHeap.Size() != 5 {
		t.Error("Peek broke the heap!")
	}

}

func Test_ExtractMinReversed(t *testing.T) {
	maxHeap := heap.NewMaxHeap()
	maxHeap.Insert(1)
	maxHeap.Insert(2)
	maxHeap.Insert(3)
	maxHeap.Insert(-3)
	maxHeap.Insert(12)

	if maxHeap.Size() != 5 {
		t.Error("Heap is the wrong size!")
	}

	min, err := maxHeap.ExtractMin()
	if min != 12 || err != nil {
		t.Error("Failed to ExtractMin!")
	}

	if maxHeap.Size() != 4 {
		t.Error("Failed to remove the minimum element!")
	}
}

func Test_PeekMax(t *testing.T) {

	maxHeap := heap.NewMaxHeap()
	maxHeap.Insert(6)
	maxHeap.Insert(-12)
	maxHeap.Insert(3)
	maxHeap.Insert(-3)
	maxHeap.Insert(12)

	if maxHeap.Size() != 5 {
		t.Error("Heap is the wrong size!")
	}

	if maxHeap.Peek() != 12 {
		t.Error("Peek() is not the minumum element!")
	}

	if maxHeap.Size() != 5 {
		t.Error("Peek broke the heap!")
	}

}
