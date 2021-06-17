package utility

import (
	"fmt"
	"testing"
)

type Node struct {
	Key string
	Count int64
}

func (n Node) Value() int64 {
	return n.Count
}

func TestNewSizedMinHeap(t *testing.T) {
	minHeap := NewSizedMinHeap(4)
	counts := []int64{8, 16, 3, 9, 10, 25, 11}
	for _, c := range counts {
		minHeap.Push(Node{
			Key:   "",
			Count: c,
		})
	}
	fmt.Println(minHeap.Heap)
}