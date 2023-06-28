package goltcd

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "case01",
			args: args{
				inorder:   []int{9, 3, 15, 20, 7},
				postorder: []int{9, 15, 7, 20, 3},
			},
			want: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
		},
		{
			name: "case02",
			args: args{
				inorder:   []int{2, 1},
				postorder: []int{2, 1},
			},
			want: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
		},
		{
			name: "case03",
			args: args{
				inorder:   []int{2, 3, 1},
				postorder: []int{3, 2, 1},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.inorder, tt.args.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
