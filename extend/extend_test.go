package extend

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseWithGenerics(t *testing.T) {
	s := []int{1, 2, 3}
	expect := []int{3, 2, 1}
	got := reverseWithGenerics(s)
	assert.Equal(t, expect, got)
}
