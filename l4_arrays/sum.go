package main

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAllTails calculates the sums of all but the first number given a collection of slices.
func SumAllTails(numbers ...[]int) []int {
	lengthOfNumbers := len(numbers)
	sum := make([]int, lengthOfNumbers)

	for i, numbers := range numbers {
		if len(numbers) == 0 {
			sum[i] = 0
		} else {
			sum[i] = Sum(numbers[1:])
		}
	}
	return sum
}
