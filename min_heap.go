package utility

import "container/heap"

type Compare interface {
	Value() int64
}

type MinHeap []Compare

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i].Value() < (*h)[j].Value()
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Compare))
}

func (h *MinHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

type SizedMinHeap struct {
	Size uint32
	Heap MinHeap
}

func NewSizedMinHeap(size uint32) *SizedMinHeap {
	h := make(MinHeap, 0, size)
	heap.Init(&h)
	return &SizedMinHeap{
		Size: size,
		Heap: h,
	}
}

func (h *SizedMinHeap) Push(x Compare) {
	heap.Push(&h.Heap, x)
	if h.Heap.Len() > int(h.Size) {
		heap.Pop(&h.Heap)
	}
}
