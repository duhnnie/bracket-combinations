package combinations

import (
	"fmt"
	"math"
)

var results = []uint64{1, 1, 2}
var stringResults = [][]string{
	{""},
	{"()"},
	{"()()", "(())"},
}

func insertAtIndex[T any](slice []T, value T, index int) []T {
	sliceLength := len(slice)
	sliceCapacity := cap(slice)
	requiredLength := int(math.Max(float64(sliceLength + 1), float64(index + 1)))

	if requiredLength > sliceCapacity {
		newSlice := make([]T, requiredLength, requiredLength * 3/2)
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[0:requiredLength]
	copy(slice[index + 1:], slice[index:])
	slice[index] = value

	return slice
}

func wrapWithBracker(str string) string {
	return fmt.Sprintf("(%s)", str)
}

func mapx[T, I string](collection []T, fn func(T) I) []I {
	results := make([]I, len(collection))

	for i, item := range collection {
		results[i] = fn(item)
	}

	return results
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

func GetCombinations(number int) []string {
	if number < len(stringResults) {
		return stringResults[number]
	}

	var combinations []string = []string{}

	for i := 0; i < number; i++ {
		// firstPartFn := func (str string) string { return fmt.Sprintf("(%s)", str) }
		firstPartCombinations := GetCombinations(i)
		firstPartCombinations = mapx(firstPartCombinations, wrapWithBracker)

		secondPartCombinations := GetCombinations(number - i - 1)

		for j := 0; j < len(firstPartCombinations); j++ {
			for k := 0; k < len(secondPartCombinations); k++ {
				combination := fmt.Sprintf("%v%v", firstPartCombinations[j], secondPartCombinations[k])
				combinations = append(combinations, combination)
			}
		}
	}

	stringResults = insertAtIndex(stringResults, combinations, number)

	return combinations
}
