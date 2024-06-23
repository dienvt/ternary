package ternary

import "testing"

func TestIf(t *testing.T) {
	tests := map[string]struct {
		cond bool
		a    interface{}
		b    interface{}
		want interface{}
	}{
		"true": {
			cond: true,
			a:    1,
			b:    2,
			want: 1,
		},
		"false": {
			cond: false,
			a:    1,
			b:    2,
			want: 2,
		},
	}
	for name, tt := range tests {
		name, tt := name, tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := If(tt.cond, tt.a, tt.b); got != tt.want {
				t.Errorf("If() = %v, want %v", got, tt.want)
			}
		})
	}
}
