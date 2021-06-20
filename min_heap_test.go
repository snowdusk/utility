package utility

import (
	"fmt"
	"sort"
	"testing"
)

type Node struct {
	Count int64
}

func (n Node) Value() int64 {
	return n.Count
}

func TestNewSizedMinHeap(t *testing.T) {
	minHeap := NewSizedMinHeap(4)
	counts := []int64{8, 16, 3, 9, 10, 25, 11, 20}
	for _, c := range counts {
		minHeap.Push(Node{c})
	}
	fmt.Println(minHeap.Heap)
	sort.Sort(&minHeap.Heap)
	fmt.Println(minHeap.Heap)
}