package sum

// Array gets array of numbers and sums it
func Array(numbers [5]int) int {
	sum := 0

	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }

	for _, num := range numbers {
		sum += num
	}

	return sum
}

// Slice gets a slice of numbers and returns its sum
func Slice(numbers []int) int {
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	return sum
}

// AllSingleSlice gets a slice of numbers and returns a slice with the sum of the nums
func AllSingleSlice(numbers []int) []int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return []int{sum}
}

// AllMultipleSlices gets multiple slices of numbers and returns a slice with the sum of the nums in each slice
func AllMultipleSlices(slices ...[]int) []int {
	// slicesCount := len(slices)
	// finalSlice := make([]int, slicesCount)

	// for index, slice := range slices {
	// 	finalSlice[index] = Slice(slice)
	// }

	var finalSlice []int

	for _, slice := range slices {
		finalSlice = append(finalSlice, Slice(slice))
	}
	return finalSlice
}

// BiggestsFromSingleSlice gets a slice of numbers and returns a slice with the biggest number
func BiggestsFromSingleSlice(numbers []int) []int {
	biggestNumber := 0

	for _, number := range numbers {
		if number >= biggestNumber {
			biggestNumber = number
		}
	}

	return []int{biggestNumber}
}

// BiggestsFromMultipleSlices gets multiple slices of numbers and returns a slice with the biggest number from each slice
func BiggestsFromMultipleSlices(slices ...[]int) []int {
	var finalSlice []int
	for _, slice := range slices {
		biggestNumber := 0

		for _, number := range slice {
			if number > biggestNumber {
				biggestNumber = number
			}
		}
		finalSlice = append(finalSlice, biggestNumber)
	}
	return finalSlice
}
