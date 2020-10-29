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

func TestSortOrderASC(t *testing.T) {
	// Arrange
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Act
	Sort(elements)

	// Assert
	assert.EqualValues(t, expected, elements, "ソート後、昇順であること") // 同じ値かつ、同じ型
}

func TestSortOrderAlreadySorted(t *testing.T) {
	// Arrange
	elements := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Act
	Sort(elements)

	// Assert
	assert.EqualValues(t, expected, elements, "ソート後、順番に変更がないこと") // 同じ値かつ、同じ型
}

// ベンチマーク
func BenchmarkBubbleSort(b *testing.B) {
	// Arrange
	elements := getElements(1000000)

	// Act
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

// ベンチマーク
func BenchmarkSort(b *testing.B) {
	// Arrange
	elements := getElements(1000000)
	// Act
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
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
