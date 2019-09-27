package gotest

import (
	"GoProject/common"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4}
	expected := 10
	actual := common.Sum(numbers)
	if actual != expected {
		t.Errorf("Expected the sum of %v to be %d but instead got %d!", numbers, expected, actual)

	}
}
