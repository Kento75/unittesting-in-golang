package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortDESC(t *testing.T) {
	// Arrange
	elements := getElements(10)
	expected := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	// Act
	Sort(elements)

	// Assert
	assert.EqualValues(t, expected, elements, "降順にソートされていること")
}

func TestSortASC(t *testing.T) {
	// Arrange
	elements := getElements(10001)
	expected := getElementsForASC(10001)
	// Act
	Sort(elements)

	// Assert
	assert.EqualValues(t, expected, elements, "昇順にソートされていること")
}

// test helper
func getElements(n int) []int {
	result := make([]int, n)
	j := 0
	for i := n; i > 0; i-- {
		result[j] = i
		j++
	}
	return result
}

func getElementsForASC(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = i + 1
	}
	return result
}
