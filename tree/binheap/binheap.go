package binheap

import (
	"github.com/kodeyeen/container/internal/cmp"
)

type BinHeap[E any] struct {
	elems []E
	cmp   cmp.Comparator[E]
}

func New[E any](cmp cmp.Comparator[E]) *BinHeap[E] {
	return &BinHeap[E]{
		elems: make([]E, 0),
		cmp:   cmp,
	}
}

func (h *BinHeap[E]) Init(elems ...E) {
	n := len(h.elems)
	h.elems = make([]E, 0, n)
	h.elems = append(h.elems, elems...)

	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
}

func (h *BinHeap[E]) Push(elem E) {
	h.elems = append(h.elems, elem)
	h.up(len(h.elems) - 1)
}

func (h *BinHeap[E]) Pop() (E, bool) {
	if len(h.elems) == 0 {
		var e E
		return e, false
	}

	n := len(h.elems) - 1

	h.swap(0, n)

	h.down(0, n)

	elem := h.elems[n-1]
	h.elems = h.elems[0 : n-1]
	return elem, true
}

func (h *BinHeap[E]) Peek() (E, bool) {
	if len(h.elems) == 0 {
		var e E
		return e, false
	}

	return h.elems[0], true
}

func (h *BinHeap[E]) less(i, j int) bool {
	return h.cmp(h.elems[i], h.elems[j]) == -1
}

func (h *BinHeap[E]) swap(i, j int) {
	h.elems[i], h.elems[j] = h.elems[j], h.elems[i]
}

func (h *BinHeap[E]) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.less(j, i) {
			break
		}
		h.swap(i, j)
		j = i
	}
}

func (h *BinHeap[E]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.less(j, i) {
			break
		}
		h.swap(i, j)
		i = j
	}
	return i > i0
}

func (h *BinHeap[E]) Clear() {
	h.elems = make([]E, 0)
}

func (h *BinHeap[E]) Len() int {
	return len(h.elems)
}