package util

import "testing"

func TestNearestPowerOf2(t *testing.T) {
	tests := []struct {
		value  int
		answer int
	}{
		{1, 1},
		{2, 2},
		{1024, 1024},
		{513, 1024},
		{45, 64},
		{3, 4},
		{2020, 2048},
	}
	for i := range tests {
		if NearestPowerOf2(tests[i].value) != tests[i].answer {
			t.Errorf("wrong answer for %d got %d expected %d", tests[i].value, NearestPowerOf2(tests[i].value), tests[i].answer)
		}
	}
}
