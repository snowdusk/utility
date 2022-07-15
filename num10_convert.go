package utility

import (
	"fmt"
	"math"
)

// 10到62进制
type Num struct {
	CharSet  []byte
	Base     int
	CharVals map[byte]int
	Zero     string
}

func NewNum(charset string) *Num {
	if charset == "" {
		panic("invalid charset")
	}
	r := &Num{}
	r.CharSet = []byte(charset)
	r.Base = len(r.CharSet)
	r.Zero = string(r.CharSet[0])
	r.CharVals = make(map[byte]int)
	for i, v := range r.CharSet {
		r.CharVals[v] = i
	}
	r.CharVals['-'] = -1 // TODO charset中不能有-
	return r
}

func (n *Num) ToStr(number int) string {
	if number == 0 {
		return n.Zero
	}
	t := number
	if number < 0 {
		number = -number
	}
	ret := make([]byte, 0, 8)
	for number > 0 {
		ret = append(ret, n.CharSet[number%n.Base])
		number /= n.Base
	}
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	if t > 0 {
		return string(ret)
	}
	return "-" + string(ret)
}

// TODO （1）负数，（2）非法字符串
func (n *Num) ToInt(number string) (int, error) {
	if number == "" {
		return 0, fmt.Errorf("empty number")
	}
	numbytes := []byte(number)
	if !n.check(numbytes) {
		return 0, fmt.Errorf("invalid number")
	}
	l, ret := len(numbytes), 0
	for i, b := range numbytes {
		ret += n.CharVals[b] * int(math.Pow(float64(n.Base), float64(l-i-1)))
	}
	return ret, nil
}

func (n *Num) check(number []byte) bool { // TODO 检查--123
	for _, b := range number {
		if _, ok := n.CharVals[b]; !ok {
			return false
		}
	}
	return true
}
