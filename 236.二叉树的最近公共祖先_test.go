package goLeetCode

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// {
		// 	name: "case01",
		// 	args: args{
		// 		root: &TreeNode{
		// 			Val: 3,
		// 			Left: &TreeNode{
		// 				Val:  5,
		// 				Left: &TreeNode{Val: 6},
		// 				Right: &TreeNode{
		// 					Val:   2,
		// 					Right: &TreeNode{Val: 7},
		// 					Left:  &TreeNode{Val: 4},
		// 				},
		// 			},
		// 			Right: &TreeNode{
		// 				Val:   1,
		// 				Left:  &TreeNode{Val: 0},
		// 				Right: &TreeNode{Val: 8},
		// 			},
		// 		},
		// 		p: &TreeNode{Val: 5},
		// 		q: &TreeNode{Val: 4},
		// 	},
		// 	want: &TreeNode{Val: 5},
		// },
		{
			name: "case02",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 3},
					},
				},
				q: &TreeNode{Val: 3},
				p: &TreeNode{Val: 2},
			},
			want: &TreeNode{Val: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
