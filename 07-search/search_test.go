package main

import "testing"

func TestSearch(t *testing.T) {
	nums := []int{10, 20, 30, 40, 50, 60, 70, 80}

	tests := []struct {
		name   string
		target int
		want   int
	}{
		{
			name:   "target exists",
			target: 60,
			want:   5,
		},
		{
			name:   "first element",
			target: 10,
			want:   0,
		},
		{
			name:   "last element",
			target: 80,
			want:   7,
		},
		{
			name:   "target not found",
			target: 90,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			linearGot := linearSearch(nums, tt.target)
			t.Logf(
				"Linear Search: target=%d, index=%d, want=%d",
				tt.target, linearGot, tt.want,
			)

			if linearGot != tt.want {
				t.Errorf("linearSearch() = %d, want %d", linearGot, tt.want)
			}

			binaryGot := binarySearch(nums, tt.target)
			t.Logf(
				"Binary Search: target=%d, index=%d, want=%d",
				tt.target, binaryGot, tt.want,
			)

			if binaryGot != tt.want {
				t.Errorf("binarySearch() = %d, want %d", binaryGot, tt.want)
			}
		})
	}
}
