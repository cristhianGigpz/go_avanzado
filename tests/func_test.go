package tests

import (
	"go-avanzado/utils"
	"testing"
)

// func TestSum(t *testing.T) {

// 	result := Sum(2, 3)

//		if result != 5 {
//			t.Errorf(
//				"expected 5 but got %d",
//				result,
//			)
//		}
//	}
func TestSum(t *testing.T) {

	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "2+3",
			a:        2,
			b:        3,
			expected: 5,
		},
		{
			name:     "5+5",
			a:        5,
			b:        5,
			expected: 10,
		},
		{
			name:     "6+8",
			a:        6,
			b:        8,
			expected: 14,
		},
	}

	for _, test := range tests {

		result := utils.Sum(test.a, test.b)

		if result != test.expected {

			t.Errorf(
				"%s expected %d got %d",
				test.name,
				test.expected,
				result,
			)
		}
	}
}
