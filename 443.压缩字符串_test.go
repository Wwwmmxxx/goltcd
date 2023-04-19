package goLeetCode

import "testing"

func Test_compress(t *testing.T) {
	type args struct {
		chars []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case01",
			args: args{
				chars: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			},
			want: 6,
		},
		{
			name: "case02",
			args: args{
				chars: []byte{'a', 'b', 'c'},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.args.chars); got != tt.want {
				t.Errorf("compress() = %v, want %v", got, tt.want)
			}
		})
	}
}
