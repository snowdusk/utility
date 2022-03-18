package utility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlag(t *testing.T) {
	test := assert.New(t)

	stringFlag := NewFlag[string]("a", "b", "c", "d")
	test.Equal("a", stringFlag.Select(""))
	test.Equal("a", stringFlag.Select("aa"))
	test.Equal("b", stringFlag.Select("b"))
	test.Equal("c", stringFlag.Select("c"))
	test.Equal("d", stringFlag.Select("d"))

	intFlag := NewFlag[int](1, 2, 3, 4)
	test.Equal(1, intFlag.Select(0))
	test.Equal(1, intFlag.Select(11))
	test.Equal(2, intFlag.Select(2))
	test.Equal(3, intFlag.Select(3))
	test.Equal(4, intFlag.Select(4))
}
