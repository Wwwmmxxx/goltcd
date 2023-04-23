package goltcd

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case01",
			args: args{
				intervals: [][]int{
					{1, 3},
					{2, 6},
					{8, 10},
					{15, 18},
				},
			},
			want: [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			name: "case02",
			args: args{
				intervals: [][]int{
					{1, 4},
					{2, 3},
				},
			},
			want: [][]int{
				{1, 4},
			},
		},
		{
			name: "case03",
			args: args{
				intervals: [][]int{
					{1, 4},
					{1, 4},
				},
			},
			want: [][]int{
				{1, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
