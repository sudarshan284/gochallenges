package main

import (
	"reflect"
	"testing"
)

func TestNumbertoMap(t *testing.T) {
	tests := []struct {
		input    int
		expected map[int]int
	}{
		{5, map[int]int{1: 1, 2: 4, 3: 9, 4: 16, 5: 25}},
		{3, map[int]int{1: 1, 2: 4, 3: 9}},
		{0, map[int]int{}}, // No values if n = 0
	}

	for _, test := range tests {
		result := NumbertoMap(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("NumbertoMap(%d) = %v; want %v", test.input, result, test.expected)
		}
	}
}
