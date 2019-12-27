package arrays_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		var expected int = 15
		var actual int = Sum(nums)

		if expected != actual {
			t.Errorf("Expected: %d, Actual: %d", expected, actual)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		var expected int = 6
		var actual int = Sum(numbers)

		if expected != actual {
			t.Errorf("Expected: %d, Actual: %d", expected, actual)
		}
	})
}

func checkSums (t *testing.T, actual, expected []int) {
  t.Helper()
  
  if !reflect.DeepEqual(actual, expected) {
    t.Errorf("actual: %v expected: %v", actual, expected)
  }
}

func TestSumAll(t *testing.T) {
	expected := []int{3, 9}
  actual := SumAll([]int{1, 2}, []int{0, 9})
  
  checkSums(t, actual, expected)
}

func TestSumAllTails(t *testing.T) {
	t.Run("create sums of slices without tail", func (t *testing.T) {
    actual := SumAllTails([]int{1, 2}, []int{0, 9})
    expected := []int{2, 9}

    checkSums(t, actual, expected)
  })

  t.Run("Safely sum empty slices and regular slices", func(t *testing.T) {
    actual := SumAllTails([]int{}, []int{1,3,9})
    expected := []int{0, 12}

    checkSums(t, actual, expected)
  })
}
