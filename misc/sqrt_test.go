package misc

import "testing"

// TestSqrtPositiveNumber calls misc.Sqrt with a positive number, checking for a valid return
// value
func TestSqrtPositiveNumber(t *testing.T) {
	result, err := Sqrt(25)
	if result != 5 {
		t.Errorf("Expected 5, found %f\nError: %s", result, err)
	}
}

func TestSqrtPositiveNumberTableDriven(t *testing.T) {
	// Define the columns
	var tests = []struct {
		name  string
		input float64
		want  float64
	}{
		{"Sqrt of 25 should be 5", 25, 5},
		{"Sqrt of 9 should be 3", 9, 3},
		{"Sqrt of 100 should be 10", 100, 10},
	}
	// Execution loop
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			result, err := Sqrt(v.input)
			if result != v.want {
				t.Errorf("Expected 5, found %f\nError: %s", result, err)
			}
		})
	}
}
