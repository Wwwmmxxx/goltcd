package goltcd

import (
	"testing"
)

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case01",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCB",
			},
			want: false,
		},
		{
			name: "case02",
			args: args{
				board: [][]byte{
					{'a', 'a', 'b', 'a', 'a', 'b'},
					{'a', 'a', 'b', 'b', 'b', 'a'},
					{'a', 'a', 'a', 'a', 'b', 'a'},
					{'b', 'a', 'b', 'b', 'a', 'b'},
					{'a', 'b', 'b', 'a', 'b', 'a'},
					{'b', 'a', 'a', 'a', 'a', 'b'},
				},
				word: "bbbaabbbbbab",
			},
			want: false,
		},
		{
			name: "case03",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'E', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCESEEEFS",
			},
			want: true,
		},
		{
			name: "case04",
			args: args{
				board: [][]byte{
					{'A', 'A', 'A', 'A', 'A', 'A'},
					{'A', 'A', 'A', 'A', 'A', 'A'},
					{'A', 'A', 'A', 'A', 'A', 'A'},
					{'A', 'A', 'A', 'A', 'A', 'A'},
					{'A', 'A', 'A', 'A', 'A', 'A'},
					{'A', 'A', 'A', 'A', 'A', 'A'},
				},
				word: "AAAAAAAAAAAABAA",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
