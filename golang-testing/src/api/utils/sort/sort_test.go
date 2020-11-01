package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// command: go test -v --cover ./golang-testing/utils/sort/

func TestBubbleSortOrderDESC(t *testing.T) {
	// Arrange
	elements := getElementsForASC(10000)
	expected := getElements(10000)

	// Act
	BubbleSort(elements)

	// Assert
	assert.EqualValues(t, expected, elements, "ソート後、降順であること") // 同じ値かつ、同じ型
}

func TestBubbleSortAlreadySorted(t *testing.T) {
	// Arrange
	elements := getElements(10000)
	expected := getElements(10000)

	// Act
	BubbleSort(elements)

	// Assert
	assert.EqualValues(t, expected, elements, "ソート後、順番に変更がないこと") // 同じ値かつ、同じ型
}

func TestSortOrderASC(t *testing.T) {
	// Arrange
	elements := getElements(10001)
	expected := getElementsForASC(10001)

	// Act
	Sort(elements)

	// Assert
	assert.EqualValues(t, expected, elements, "ソート後、昇順であること") // 同じ値かつ、同じ型
}

func TestSortOrderAlreadySorted(t *testing.T) {
	// Arrange
	elements := getElementsForASC(10001)
	expected := getElementsForASC(10001)

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

func getElementsForASC(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = i + 1
	}
	return result
}
