package main

import "testing"

func Test_reverse2(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				x: 123,
			},
			want: 321,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse2(tt.args.x); got != tt.want {
				t.Errorf("reverse2() = %v, want %v", got, tt.want)
			}
		})
	}
}
