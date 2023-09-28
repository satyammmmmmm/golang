package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	ele := []int{9, 8, 7, 6, 5}
	ele = BubbleSort(ele)
	assert.NotNil(t, ele)
	assert.EqualValues(t, 5, len(ele))
	assert.EqualValues(t, 5, ele[0])
	assert.EqualValues(t, 6, ele[1])

}
func getElement(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func BenchmarkBubbleSort(b *testing.B) {
	ele := getElement((10))
	for i := 0; i < b.N; i++ {
		BubbleSort(ele)
	}
}
