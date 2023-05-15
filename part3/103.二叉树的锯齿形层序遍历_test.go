package goltcd

import (
	"reflect"
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case01",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 4},
					},
					Right: &TreeNode{
						Val:   3,
						Right: &TreeNode{Val: 5},
					},
				},
			},
			want: [][]int{
				{1},
				{3, 2},
				{4, 5},
			},
		}, {
			name: "case02",
			args: args{
				root: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:   1,
							Left:  &TreeNode{Val: 5},
							Right: &TreeNode{Val: 1},
						},
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:   3,
							Right: &TreeNode{Val: 6},
						},
						Right: &TreeNode{
							Val:   -1,
							Right: &TreeNode{Val: 8},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigzagLevelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
