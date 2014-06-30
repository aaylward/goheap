package heap

type MinHeap struct {
	list []int
	reversed bool
}

func (h *MinHeap) Less(i, j int) bool {
	if h.reversed {
		return h.list[i] > h.list[j]
	}
	return h.list[i] < h.list[j]
}

func (h *MinHeap) Insert(x int) {
	idx := h.Size()
	h.list = append(h.list, x)
	up(h, idx)
}

func (h *MinHeap) ExtractMin() int {
	idx := h.Size() - 1
	swap(h, 0, idx)
	x := h.list[idx]
	h.list = h.list[:idx]
	down(h)
	return x
}

func (h MinHeap) Size() int {
	return len(h.list)
}

func (h MinHeap) Peek() int {
	return h.list[0]
}

func up (h *MinHeap, child int) {
	var parent int
	for {
		if child%2==0 {
			parent = (child-1)/2
		} else {
			parent = child/2
		}
		if h.Less(child, parent) {
			swap(h, child, parent)
			child = parent
		} else {
			break
		}
	}
}

func down(h *MinHeap) {
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

func swap(h *MinHeap, i, j int) {
	h.list[i], h.list[j] = h.list[j], h.list[i]
}

