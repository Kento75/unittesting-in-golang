package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// command: go test -v --cover ./golang-testing/utils/sort/

func TestBubbleSortOrderDESC(t *testing.T) {
	// Arrange
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	expected := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	// Act
	BubbleSort(elements)

	// Assert
	assert.EqualValues(t, expected, elements, "ソート後、降順であること") // 同じ値かつ、同じ型
}

func TestBubbleSortAlreadySorted(t *testing.T) {
	// Arrange
	elements := []int{5, 4, 3, 2, 1, 0}
	expected := []int{5, 4, 3, 2, 1, 0}

	// Act
	BubbleSort(elements)

	// Assert
	assert.EqualValues(t, expected, elements, "ソート後、順番に変更がないこと") // 同じ値かつ、同じ型
}
