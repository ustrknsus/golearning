package main

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct{ a, b, want int }{{2, 3, 5}, {-2, 2, 0}, {0, 0, 0}}
	for _, tc := range cases {
		if got := Add(tc.a, tc.b); got != tc.want {
			t.Errorf("Add(%d,%d)=%d; want %d", tc.a, tc.b, got, tc.want)
		}
	}
}
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Add(i, 1)
	}
}
