package utility

import (
	"fmt"
	"testing"
)

func TestStatisBucket(t *testing.T) {
	b := NewStatisBucket(5, 5)
	b.Incr(40, 4)
	b.Incr(15, 3)
	b.Incr(10, 1)
	b.Incr(100, 99)
	b.Incr(1, 1)
	b.Incr(3, 44)
	ret := b.StatisInfo(1)
	for _, r := range ret {
		fmt.Println(r)
	}
}
