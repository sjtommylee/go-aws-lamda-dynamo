package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		// Add more test cases as needed
	}

	for _, test := range tests {
		solution := TwoSum{nums: test.nums, target: test.target}
		result := solution.hashmap()
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For nums %v and target %d, expected %v, but got %v", test.nums, test.target, test.expected, result)
		}
	}
}
