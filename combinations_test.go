package combinations

import "testing"

var inputs [10]int = [10]int { 0, 1, 2, 3, 4, 5, 7, 20, 25, 30 }
var outputs [10]uint64 = [10]uint64 { 1, 1, 2, 5, 14, 42, 429, 6564120420, 4861946401452, 3814986502092304 }

var stringOutputs [][]string = [][]string {
	{""},
	{
		"()",
	},
	{
		"()()",
		"(())",
	},
	{
		"()()()",
		"()(())",
		"(())()",
		"(()())",
		"((()))",
	},
	{
		"()()()()",
		"()()(())",
		"()(())()",
		"()(()())",
		"()((()))",
		"(())()()",
		"(())(())",
		"(()())()",
		"((()))()",
		"(()()())",
		"(()(()))",
		"((())())",
		"((()()))",
		"(((())))",
	},
	{
		"()()()()()",
		"()()()(())",
		"()()(())()",
		"()()(()())",
		"()()((()))",
		"()(())()()",
		"()(())(())",
		"()(()())()",
		"()((()))()",
		"()(()()())",
		"()(()(()))",
		"()((())())",
		"()((()()))",
		"()(((())))",
		"(())()()()",
		"(())()(())",
		"(())(())()",
		"(())(()())",
		"(())((()))",
		"(()())()()",
		"(()())(())",
		"((()))()()",
		"((()))(())",
		"(()()())()",
		"(()(()))()",
		"((())())()",
		"((()()))()",
		"(((())))()",
		"(()()()())",
		"(()()(()))",
		"(()(())())",
		"(()(()()))",
		"(()((())))",
		"((())()())",
		"((())(()))",
		"((()())())",
		"(((()))())",
		"((()()()))",
		"((()(())))",
		"(((())()))",
		"(((()())))",
		"((((()))))",
	},
}

func TestGetCombinationsCount(t *testing.T) {
	for i := 0; i < len(inputs); i++ {
		output := GetCombinationsCount(inputs[i])

		if output != outputs[i] {
			t.Fatalf("GetCombinationsCount(%v) = %v. %v was expected.", inputs[i], output, outputs[i])
		}
	}
}

func TestGetCombinations(t *testing.T) {
	// Here we only test until input 5,
	// this is because this function is expensive
	for i := 0; i < 6; i++ {
		output := GetCombinations(i)
		outputLength := len(output)
		expectedLength := len(stringOutputs[i])

		if outputLength != expectedLength {
			t.Fatalf("Incorrect length for GetCombinations(%v) = %v. Expected: %v", i, outputLength, expectedLength)
		}

		for j := 0; j < outputLength; j++ {
			if output[j] != stringOutputs[i][j] {
				t.Fatalf("Wrong output for GetCombinations(%v) = %v. Expected: %v", i, output[j], stringOutputs[i][j])
			}
		}
	}
}
