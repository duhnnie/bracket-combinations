package combinations

import "testing"

func TestGetCombinationsCount(t *testing.T) {
	var inputs [10]int = [10]int { 0, 1, 2, 3, 4, 5, 7, 20, 25, 30 }
	var outputs [10]uint64 = [10]uint64 { 1, 1, 2, 5, 14, 42, 429, 6564120420, 4861946401452, 3814986502092304 }

	for i := 0; i < len(inputs); i++ {
		output := GetCombinationsCount(inputs[i])

		if output != outputs[i] {
			t.Fatalf("GetCombinationsCount(%v) = %v. %v was expected.", inputs[i], output, outputs[i])
		}
	}
}
