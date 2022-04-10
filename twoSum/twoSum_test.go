package twoSum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	nums   = []int{2, 7, 11, 15}
	target = 9
)

func TestTwoSumFunc1(t *testing.T) {
	result := TwoSum1(nums, target)
	expected := []int{2, 7}
	assert.Equal(t, expected, result)
}

func TestTwoSumFunc2(t *testing.T) {
	result := TwoSum2(nums, target)
	expected := []int{2, 7}
	assert.Equal(t, expected, result)
}
