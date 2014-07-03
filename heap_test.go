package heap_test

import (
	"github.com/aaylward/heap"
	"testing"
)

func Test_InsertToHeap(t *testing.T) {
	minHeap := heap.NewHeap(false)
	minHeap.Insert(5)

	if minHeap.Size() != 1 {
		t.Error("Failed to insert an element!")
	} else {
		t.Log("Successfully inserted an element!")
	}
}

func Test_ExtractMin(t *testing.T) {
	minHeap := heap.NewHeap(false)
	minHeap.Insert(1)
	minHeap.Insert(2)
	minHeap.Insert(3)
	minHeap.Insert(-3)
	minHeap.Insert(12)

	if minHeap.Size() != 5 {
		t.Error("Heap is the wrong size!")
	}

	if minHeap.ExtractMin() != -3 {
		t.Error("Failed to ExtractMin!")
	}

	if minHeap.Size() != 4 {
		t.Error("Failed to remove the minimum element!")
	}
}

func Test_Peek(t *testing.T) {

	minHeap := heap.NewHeap(false)
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
