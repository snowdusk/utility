package utility

import "fmt"

type StatisBucket struct {
	BucketSize int64 // 每个桶的大小
	BucketNums int64 // 桶的数量
	Buckets    []int64
}

func NewStatisBucket(size, nums int64) *StatisBucket {
	return &StatisBucket{
		BucketSize: size,
		BucketNums: nums,
		Buckets:    make([]int64, nums),
	}
}

type Statis struct {
	Range string
	Sum   int64
}

func (s *StatisBucket) StatisInfo(step int64) []*Statis {
	var stepCount int64
	if s.BucketNums%step == 0 {
		stepCount = s.BucketNums / step
	} else {
		stepCount = s.BucketNums/step + 1
	}
	ret := make([]*Statis, 0, stepCount)
	for i := int64(0); i < stepCount; i++ {
		start, end := i*step, (i+1)*step
		if end > s.BucketNums {
			end = s.BucketNums
		}
		sum := int64(0)
		for _, c := range s.Buckets[start:end] {
			sum += c
		}
		var endStr interface{} = end * s.BucketSize
		if end == s.BucketNums {
			endStr = "+∞"
		}
		ret = append(ret, &Statis{fmt.Sprintf("[%d, %v):", start*s.BucketSize, endStr), sum})
	}
	return ret
}

func (s *StatisBucket) Incr(count int64, delta int64) {
	bucketIdx := count / s.BucketSize
	if bucketIdx > s.BucketNums-1 {
		bucketIdx = s.BucketNums - 1
	}
	s.Buckets[bucketIdx] += delta
}
