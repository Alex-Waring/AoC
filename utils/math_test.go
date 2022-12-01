package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	type test struct {
		input []int
		want  int
	}

	tests := []test{
		{input: []int{1, 2, 3}, want: 6},
		{input: []int{}, want: 0},
	}

	for _, tc := range tests {
		got := Sum(tc.input)
		assert.Equal(t, tc.want, got)
	}
}

func TestMakeRange(t *testing.T) {
	type input struct {
		min int
		max int
	}
	type test struct {
		input input
		want  []int
	}

	tests := []test{
		{input: input{min: 0, max: 0}, want: []int{0}},
		{input: input{min: 1, max: 7}, want: []int{1, 2, 3, 4, 5, 6, 7}},
	}

	for _, tc := range tests {
		got := MakeRange(tc.input.min, tc.input.max)
		assert.Equal(t, tc.want, got)
	}
}
