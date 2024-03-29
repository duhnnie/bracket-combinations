package combinations

import "math"

var results = []uint64{1, 1, 2}

func insertAtIndex(slice []uint64, number uint64, index int) []uint64 {
	sliceLength := len(slice)
	sliceCapacity := cap(slice)
	requiredLength := int(math.Max(float64(sliceLength + 1), float64(index + 1)))

	if requiredLength > sliceCapacity {
		var newSlice = make([]uint64, requiredLength, requiredLength * 3/2)
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[0:requiredLength]
	copy(slice[index + 1:], slice[index:])
	slice[index] = number

	return slice
}

// It receives a number of brackets pairs
// and returns the number of all possible combinations
func GetCombinationsCount(number int) uint64 {
	if number < len(results) {
		return results[number]
	}

	var sum uint64 = 0

	for i := 1; i <= number; i ++ {
		sum += (GetCombinationsCount(number - i) * GetCombinationsCount(i - 1))
	}

	results = insertAtIndex(results, sum, number)

	return sum
}
