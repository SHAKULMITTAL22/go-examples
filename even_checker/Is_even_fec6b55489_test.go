package main

import "testing"

func TestIs_even_fec6b55489(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{"positive even number", 2, 1},
		{"negative even number", -2, 1},
		{"positive odd number", 1, 0},
		{"negative odd number", -1, 0},
		{"zero", 0, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := is_even(tt.num); got != tt.want {
				t.Errorf("is_even(%d) = %d, want %d", tt.num, got, tt.want)
			}
		})
	}
}
