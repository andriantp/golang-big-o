package main

import "testing"

func TestLookup(t *testing.T) {
	sliceItems := []string{"camera-1", "camera-2", "camera-3"}

	mapItems := make(map[string]struct{}, len(sliceItems))
	for _, item := range sliceItems {
		mapItems[item] = struct{}{}
	}

	tests := []struct {
		name   string
		target string
		want   bool
	}{
		{
			name:   "target exists",
			target: "camera-2",
			want:   true,
		},
		{
			name:   "target not found",
			target: "camera-4",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sliceGot := findInSlice(sliceItems, tt.target)
			t.Logf(
				"Slice lookup: target=%s, found=%t, want=%t",
				tt.target, sliceGot, tt.want,
			)

			if sliceGot != tt.want {
				t.Errorf("findInSlice() = %t, want %t", sliceGot, tt.want)
			}

			mapGot := findInMap(mapItems, tt.target)
			t.Logf(
				"Map lookup: target=%s, found=%t, want=%t",
				tt.target, mapGot, tt.want,
			)

			if mapGot != tt.want {
				t.Errorf("findInMap() = %t, want %t", mapGot, tt.want)
			}
		})
	}
}
