package number

import "testing"

func TestIsNumberMatchLength(t *testing.T) {
	type args struct {
		n      int
		length int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Scenario1", args{100, 2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumberMatchLength(tt.args.n, tt.args.length); got != tt.want {
				t.Errorf("IsNumberMatchLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
