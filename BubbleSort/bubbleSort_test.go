package bubblesort

import "testing"

func TestSort(t *testing.T) {
	var tests = []struct {
		input    [3]int
		expected [3]int
	}{
		{[...]int{1, 2, 3}, [...]int{1, 2, 3}},
	}

	for _, test := range tests {
		if output := BubbleSort(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}")
		}
	}
}
