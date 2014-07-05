package heap

import "fmt"

func NewMinHeap() *Heap {
	return &Heap{reversed: false}
}

func NewMaxHeap() *Heap {
	return &Heap{reversed: true}
}

type Heap struct {
	list     []int
	reversed bool
}

func (h *Heap) Less(i, j int) bool {
	if h.reversed {
		return h.list[i] > h.list[j]
	}
	return h.list[i] < h.list[j]
}

func (h *Heap) Insert(x int) {
	idx := h.Size()
	h.list = append(h.list, x)
	up(h, idx)
}

func (h *Heap) ExtractMin() (int, error) {
  if h.Size() < 1 {
    return 0, fmt.Errorf("Cannot ExtractMin from an empty heap.")
  }

	idx := h.Size() - 1
	swap(h, 0, idx)
	x := h.list[idx]
	h.list = h.list[:idx]
	down(h)
	return x, nil
}

func (h Heap) Size() int {
	return len(h.list)
}

func (h Heap) Peek() int {
	return h.list[0]
}

func up(h *Heap, child int) {
	var parent int
	for {
		if child%2 == 0 {
			parent = (child - 1) / 2
		} else {
			parent = child / 2
		}
		if h.Less(child, parent) {
			swap(h, child, parent)
			child = parent
		} else {
			break
		}
	}
}

func down(h *Heap) {
	var n int
	for {
		left := 2*n + 1
		right := left + 1
		if right < h.Size() {
			if h.Less(left, right) {
				swap(h, n, left)
				n = left
			} else {
				swap(h, n, right)
				n = right
			}
		} else if left < h.Size() && h.Less(left, n) {
			swap(h, n, left)
			n = left
		} else {
			break
		}
	}
}

func swap(h *Heap, i, j int) {
	h.list[i], h.list[j] = h.list[j], h.list[i]
}
