package arrays_slices

// Sum func returns the sum of the integers in the array.
func Sum(nums []int) int {
	var sum int = 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

// SumAll func sums every provided slice. returns []int where each element is the sum of each slice
func SumAll(slicesToSum ...[]int) []int {
  var sums []int

  for _, numbers := range slicesToSum {
    sums = append(sums, Sum(numbers))
  }

	return sums
}

// SumAllTails func similar to SumAll except doesn't include head in sum and empty slices.
func SumAllTails(slicesToSum ...[]int) []int {
  var sums []int

  for _, numbers := range slicesToSum {
    var sum int = 0

    if len(numbers) > 0 {
      sum = Sum(numbers[1:])
    }
    
    sums = append(sums, sum)
  }

  return sums
}
