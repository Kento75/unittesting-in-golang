package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSortOrderDESC(t *testing.T) {
	// Arrange
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	expected := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	// Act
	BubbleSort(elements)

	// Assert
	assert.EqualValues(t, expected, elements) // 同じ値かつ、同じ型
}
